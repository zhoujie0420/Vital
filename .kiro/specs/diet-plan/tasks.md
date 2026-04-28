# Implementation Plan: Diet Plan (饮食规划)

## Overview

实现饮食规划功能，允许用户通过预设模板或自定义方案设定每日营养目标，并在首页展示当日饮食进度。采用纯前端方案，数据通过 localStorage 持久化。按优先级递进实现：P0 自定义方案 + 首页模块 → P1 固定模板自动计算 → P2 碳循环 → P3 碳水渐降。

## Tasks

- [x] 1. Set up project structure, install dependencies, and register pages
  - Install `fast-check` as a devDependency for property-based testing (`npm install --save-dev fast-check`)
  - Install `vitest` as a devDependency for running tests (`npm install --save-dev vitest`)
  - Create empty files: `src/utils/dietPlanTemplates.js`, `src/utils/dietPlanCalc.js`, `src/store/dietPlan.js`
  - Create page files: `src/pages/diet/plan.vue`, `src/pages/diet/plan-edit.vue`
  - Register `pages/diet/plan` and `pages/diet/plan-edit` in `src/pages.json` with `navigationStyle: "custom"`
  - Add a `"test"` script to `package.json`: `"test": "vitest --run"`
  - _Requirements: 10.1_

- [x] 2. Implement public templates and calculation logic
  - [x] 2.1 Create public template definitions in `dietPlanTemplates.js`
    - Define the 5 hardcoded `PUBLIC_TEMPLATES` array: 均衡饮食 (fixed), 碳循环 (carb_cycle), 碳水渐降 (carb_taper), 高蛋白增肌 (fixed), 低碳减脂 (fixed)
    - Each template includes: id, scope='public', name, description, type, calcConfig
    - Include `defaultCycleConfig` for carb_cycle template and `defaultTaperConfig` for carb_taper template
    - _Requirements: 1.8_

  - [x] 2.2 Implement `calculateMacrosFromTemplate(template, bodyWeight)` in `dietPlanCalc.js`
    - For ratio-based templates (均衡饮食): calculate calories = bodyWeight * caloriesPerKg, then derive protein/carbs/fat from ratios
    - For per-kg templates (高蛋白增肌, 低碳减脂): calculate protein = bodyWeight * proteinPerKg, derive remaining from ratios
    - Return `{ calories, protein, carbs, fat }` with all non-negative values
    - _Requirements: 2.1, 2.2, 2.3, 2.4_

  - [x] 2.3 Implement `getTodayTargets(plan, date)` dispatcher in `dietPlanCalc.js`
    - For `fixed` type: return `plan.macroTargets` directly
    - For `carb_cycle` type: delegate to `getCarbCycleDayTargets(plan, date)`
    - For `carb_taper` type: delegate to `getCarbTaperDayTargets(plan, date)`
    - _Requirements: 5.3, 6.3, 7.6, 7.7_

  - [x] 2.4 Implement `getCarbCycleDayTargets(plan, date)` in `dietPlanCalc.js`
    - Get day of week from date (0=Sunday..6=Saturday)
    - Look up `plan.cycleConfig.dayMap[dayOfWeek]` to determine 'high' or 'low'
    - Return corresponding `highDayTargets` or `lowDayTargets`
    - _Requirements: 5.3, 7.6_

  - [x] 2.5 Implement `getCarbTaperDayTargets(plan, date)` in `dietPlanCalc.js`
    - Calculate weekNumber = floor((date - startedAt) / 7 days) + 1
    - Calculate current carbs = max(initialCarbs - (weekNumber - 1) * weeklyReduction, minCarbs)
    - Cap carbs at minimum 50g threshold
    - Return adjusted macroTargets
    - _Requirements: 6.2, 6.3, 6.4, 7.7_

  - [x] 2.6 Implement `calculateRemaining(targets, consumed)` in `dietPlanCalc.js`
    - Return `{ calories: targets.calories - consumed.calories, protein: ..., carbs: ..., fat: ... }`
    - _Requirements: 7.1_

  - [x] 2.7 Implement validation functions in `dietPlanCalc.js`
    - `validatePlanInput(data)`: calories > 0, protein/carbs/fat >= 0
    - `validateCycleConfig(dayMap)`: at least one 'high' and one 'low' day
    - `validateTaperConfig(config)`: totalWeeks is positive integer, weeklyReduction is positive number
    - Return `{ valid, errors }` object with specific error messages per field
    - _Requirements: 3.2, 3.3, 3.5, 5.5, 6.5_

  - [x] 2.8 Write property test: Template aggregation returns all templates (Property 1)
    - **Property 1: Template aggregation returns all templates**
    - **Validates: Requirements 1.1**

  - [x] 2.9 Write property test: Macro calculation from fixed template (Property 4)
    - **Property 4: Macro calculation from fixed template**
    - **Validates: Requirements 2.1, 2.2, 2.3, 2.4**

  - [x] 2.10 Write property test: Plan numeric input validation (Property 6)
    - **Property 6: Plan numeric input validation**
    - **Validates: Requirements 3.2, 3.3, 6.5**

  - [x] 2.11 Write property test: Carb cycle day-of-week target lookup (Property 9)
    - **Property 9: Carb cycle day-of-week target lookup**
    - **Validates: Requirements 5.3, 7.6**

  - [x] 2.12 Write property test: Cycle config requires mixed day types (Property 10)
    - **Property 10: Cycle config requires mixed day types**
    - **Validates: Requirements 5.5**

  - [x] 2.13 Write property test: Carb taper weekly target calculation (Property 11)
    - **Property 11: Carb taper weekly target calculation**
    - **Validates: Requirements 6.2, 6.3, 6.4, 7.7**

  - [x] 2.14 Write property test: Remaining calories calculation (Property 12)
    - **Property 12: Remaining calories calculation**
    - **Validates: Requirements 7.1**

  - [x] 2.15 Write unit tests for specific template calculations
    - Test 均衡饮食 template with bodyWeight=70kg produces expected macro values
    - Test 高蛋白增肌 template calculates protein as 2.0g/kg
    - Test 低碳减脂 template produces reduced carb ratio
    - Test bodyWeight=0 or null triggers appropriate error
    - _Requirements: 2.2, 2.3, 2.4, 2.7_

- [x] 3. Checkpoint - Ensure all calculation and validation tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 4. Implement Pinia store `useDietPlanStore`
  - [x] 4.1 Create `store/dietPlan.js` with state, getters, and localStorage persistence
    - State: `plans` (UserDietPlan[]), `privateTemplates` (DietPlanTemplate[]), `activePlanId` (string|null)
    - Getters: `activePlan`, `allTemplates` (public + private), `publicTemplates`
    - Implement `_persist()` to save plans, privateTemplates, activePlanId to localStorage
    - Implement `_restore()` to load from localStorage with JSON.parse error handling
    - Implement `init()` action that calls `_restore()` on store initialization
    - Follow existing `useUserStore` Options API pattern with `defineStore`
    - _Requirements: 8.1, 8.2, 8.3, 8.4, 8.5_

  - [x] 4.2 Implement plan CRUD actions in the store
    - `createPlanFromTemplate(templateId, overrides, bodyWeight)`: find template, calculate macros via `calculateMacrosFromTemplate`, copy params as snapshot, generate UUID, persist
    - `createCustomPlan(planData)`: validate input via `validatePlanInput`, create plan with type='fixed', persist
    - `updatePlan(planId, updates)`: find plan, merge updates preserving id/userId/createdAt, persist
    - `deletePlan(planId)`: remove plan, if it was active set activePlanId to null, persist
    - _Requirements: 2.5, 2.6, 3.4, 9.1, 9.2, 9.3_

  - [x] 4.3 Implement activation/deactivation actions in the store
    - `activatePlan(planId)`: set all plans' isActive to false, then set target plan's isActive to true, update activePlanId, persist
    - `deactivatePlan(planId)`: set plan's isActive to false, set activePlanId to null, persist
    - Enforce single-active invariant: at most one plan active at any time
    - _Requirements: 4.1, 4.2, 4.3_

  - [x] 4.4 Implement private template CRUD actions in the store
    - `createPrivateTemplate(templateData)`: set scope='private', userId=current user, generate UUID, persist
    - `updatePrivateTemplate(templateId, updates)`: find template, reject if scope='public', merge updates, persist
    - `deletePrivateTemplate(templateId)`: find template, reject if scope='public', remove, persist
    - _Requirements: 1.4, 1.5, 1.6, 1.7_

  - [-] 4.5 Write property test: Private template creation sets correct metadata (Property 2)
    - **Property 2: Private template creation sets correct metadata**
    - **Validates: Requirements 1.4**

  - [~] 4.6 Write property test: Public templates are immutable (Property 3)
    - **Property 3: Public templates are immutable**
    - **Validates: Requirements 1.7**

  - [~] 4.7 Write property test: Template-to-plan snapshot isolation (Property 5)
    - **Property 5: Template-to-plan snapshot isolation**
    - **Validates: Requirements 2.5, 2.6**

  - [~] 4.8 Write property test: Custom plan creation produces fixed-type plan (Property 7)
    - **Property 7: Custom plan creation produces fixed-type plan**
    - **Validates: Requirements 3.4**

  - [~] 4.9 Write property test: Single active plan invariant (Property 8)
    - **Property 8: Single active plan invariant**
    - **Validates: Requirements 4.1, 4.2**

  - [~] 4.10 Write property test: localStorage persistence round-trip (Property 13)
    - **Property 13: localStorage persistence round-trip**
    - **Validates: Requirements 8.1, 8.2, 8.3**

  - [~] 4.11 Write property test: In-place plan update preserves identity (Property 14)
    - **Property 14: In-place plan update preserves identity**
    - **Validates: Requirements 9.1**

  - [~] 4.12 Write property test: Deleting active plan clears active state (Property 15)
    - **Property 15: Deleting active plan clears active state**
    - **Validates: Requirements 9.3**

  - [~] 4.13 Write unit tests for store CRUD operations
    - Test createPlanFromTemplate copies params as independent snapshot
    - Test deletePlan removes plan from list and updates localStorage
    - Test activatePlan deactivates previously active plan
    - Test store init restores data from localStorage correctly
    - Test corrupted localStorage data is handled gracefully (reset to empty)
    - _Requirements: 2.5, 4.2, 8.3, 9.2_

- [~] 5. Checkpoint - Ensure all store tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 6. Implement Plan List Page (`diet/plan.vue`)
  - [~] 6.1 Build the Plan List Page layout and active plan display
    - Custom nav bar with page title "饮食规划"
    - Active plan section at top: highlight card showing current active plan name, type, and macro targets
    - If no active plan, show empty state with guidance text and "创建方案" button
    - Use SCSS with `@import '../../styles/variables.scss'` and existing mixins (`@include card`, `@include page-title`, etc.)
    - _Requirements: 4.4, 4.5_

  - [~] 6.2 Build the user plans list section
    - List all user plans with name, type badge, and macro summary
    - Each plan card has: activate/deactivate toggle, edit button, delete button
    - Delete action uses `uni.showModal` for confirmation
    - Activate action calls `useDietPlanStore.activatePlan(planId)`
    - _Requirements: 4.2, 4.3, 9.2, 9.4_

  - [~] 6.3 Build the templates section
    - List all templates (public + private) with name, description, type, and scope label
    - Visually distinguish public vs private templates (e.g., "系统" vs "自定义" badge)
    - Public templates: tap to navigate to plan-edit with `?mode=template&templateId=xxx`
    - Private templates: show edit/delete actions; delete uses `uni.showModal` confirmation
    - "创建自定义方案" button navigates to `?mode=custom`
    - _Requirements: 1.1, 1.2, 1.3, 1.5, 1.6, 1.7_

- [ ] 7. Implement Plan Edit Page (`diet/plan-edit.vue`)
  - [~] 7.1 Build the Plan Edit Page with mode routing
    - Parse URL params: `mode` (template/custom/edit), `templateId`, `planId`
    - Custom nav bar with dynamic title: "从模板创建" / "自定义方案" / "编辑方案"
    - On load: if mode=template, load template and auto-calculate macros; if mode=edit, load existing plan; if mode=custom, show empty form
    - _Requirements: 10.3_

  - [~] 7.2 Implement fixed-type plan form
    - Input fields: plan name, daily calories (kcal), protein (g), carbs (g), fat (g)
    - For template mode: pre-fill with auto-calculated values from `calculateMacrosFromTemplate`
    - If user's body weight is not set, show prompt to enter weight before auto-calculation
    - Allow manual adjustment of all auto-calculated values
    - Real-time validation with error messages under each field
    - _Requirements: 2.7, 2.8, 3.1, 3.2, 3.3, 3.5_

  - [~] 7.3 Implement carb cycle configuration form
    - Weekly calendar UI: 7 day buttons (周日-周六), tap to toggle between high/low
    - Separate macro input sections for high-carb day targets and low-carb day targets
    - Validate at least one high and one low day before submission
    - Pre-fill with template's `defaultCycleConfig` and auto-calculated values
    - _Requirements: 5.1, 5.2, 5.4, 5.5_

  - [~] 7.4 Implement carb taper configuration form
    - Input fields: total weeks (positive integer), weekly carb reduction (grams)
    - Display initial carb target (auto-calculated or manual)
    - Show preview of week-by-week carb targets with 50g minimum cap indicator
    - Validate totalWeeks and weeklyReduction
    - _Requirements: 6.1, 6.5_

  - [~] 7.5 Implement form submission and navigation
    - On submit: validate all inputs, call appropriate store action (createPlanFromTemplate / createCustomPlan / updatePlan)
    - On success: show toast "保存成功", navigate back to plan list page
    - On validation failure: display specific error messages, prevent submission
    - _Requirements: 3.4, 3.5, 9.1, 10.4_

- [ ] 8. Implement Homepage Diet Module (update `index/index.vue`)
  - [~] 8.1 Add diet plan module to homepage with active plan display
    - Import `useDietPlanStore` and call `init()` on page show
    - When active plan exists: display remaining calories (target - consumed), consumed calories, protein/carbs/fat progress bars
    - Fetch today's diet records from existing `getDietRecords` API to calculate consumed totals
    - Use `getTodayTargets(activePlan, new Date())` for current day's targets (handles fixed/carb_cycle/carb_taper)
    - _Requirements: 7.1, 7.2, 7.6, 7.7_

  - [~] 8.2 Add quick-entry buttons and plan settings link
    - Four meal-type quick buttons: 早餐, 午餐, 晚餐, 加餐 — each navigates to `/pages/diet/add?meal=xxx`
    - Plan settings link navigates to `/pages/diet/plan`
    - _Requirements: 7.3, 7.4, 10.2_

  - [~] 8.3 Implement no-plan guidance state
    - When no active plan: show guidance card with message "设置饮食规划，追踪每日营养目标"
    - Include "去设置" button linking to `/pages/diet/plan`
    - _Requirements: 7.5_

  - [~] 8.4 Implement overage warning display
    - When remaining calories < 0: display overage amount with warning color (red)
    - Use `$color-red` from design system for visual warning indicator
    - _Requirements: 7.8_

- [~] 9. Final checkpoint - Ensure all tests pass and pages are wired together
  - Ensure all tests pass, ask the user if questions arise.
  - Verify page navigation: homepage → plan list → plan edit → back to plan list
  - Verify plan list → homepage diet module reflects active plan changes

## Notes

- Tasks marked with `*` are optional and can be skipped for faster MVP
- Each task references specific requirements for traceability
- Checkpoints ensure incremental validation
- Property tests validate universal correctness properties from the design document
- Unit tests validate specific examples and edge cases
- All data persisted via localStorage; no backend API changes needed for this feature
- The existing `getDietRecords` API is reused for consumed calorie calculations on the homepage
