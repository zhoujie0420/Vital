# Requirements Document

## Introduction

Vital Fitness 是一款基于 uni-app Vue 3 的微信小程序健身追踪应用。目前已有饮食记录功能（记录每餐食物、热量、宏量营养素）。本需求文档定义「饮食规划」功能——允许用户设定每日营养目标（热量、蛋白质、碳水化合物、脂肪），通过选择预设模板或自定义方案来指导日常饮食。

系统提供两类模板：公共模板（系统预设，所有用户可见，不可编辑）和自定义模板（用户创建，仅创建者可见，可编辑删除）。模板类型包括固定比例（fixed）、碳循环（carb_cycle）和碳水渐降（carb_taper）。用户选择模板后，参数被复制到个人方案中，后续模板变更不影响已有方案。

## Glossary

- **Diet_Plan_System**: 饮食规划系统，负责管理饮食方案模板和用户饮食方案的完整子系统
- **Template**: 饮食方案模板，定义每日宏量营养素目标的配置蓝图，分为公共模板和自定义模板
- **Public_Template**: 公共模板，系统预设的饮食方案模板，所有用户可见，用户不可编辑或删除，初期在前端硬编码
- **Private_Template**: 自定义模板，用户自行创建的饮食方案模板，仅创建者可见，可编辑和删除
- **User_Diet_Plan**: 用户饮食方案，用户当前使用的饮食方案实例，包含具体的每日营养目标参数
- **Macro_Targets**: 宏量营养素目标，包括每日热量（千卡）、蛋白质（克）、碳水化合物（克）、脂肪（克）的目标值
- **Fixed_Plan**: 固定比例方案，每天使用相同的宏量营养素目标
- **Carb_Cycle_Plan**: 碳循环方案，按星期几区分高碳日（训练日）和低碳日（休息日），不同日期使用不同的碳水目标
- **Carb_Taper_Plan**: 碳水渐降方案，按周递减碳水摄入量，当前周数根据方案开始日期计算
- **Cycle_Config**: 碳循环配置，定义一周七天中每天是高碳日还是低碳日，以及对应的宏量营养素目标
- **Taper_Config**: 碳水渐降配置，定义总周数和每周的碳水递减量
- **Body_Weight**: 用户体重，存储在用户档案（User.Weight 字段）中，用于自动计算模板参数
- **Homepage_Diet_Module**: 首页饮食模块，嵌入在首页中展示当日饮食目标完成情况的 UI 组件
- **Remaining_Calories**: 剩余可摄入热量，计算公式为：目标热量 - 已摄入热量
- **Plan_List_Page**: 方案列表页，展示当前方案、切换方案、创建新方案的页面
- **Plan_Edit_Page**: 方案编辑页，选择模板或自定义配置方案参数的页面

## Requirements

### Requirement 1: Template Management

**User Story:** As a user, I want to browse and manage diet plan templates, so that I can choose a suitable nutrition strategy for my goals.

#### Acceptance Criteria

1. WHEN the Plan_List_Page loads, THE Diet_Plan_System SHALL display all Public_Template entries and all Private_Template entries belonging to the current user
2. THE Diet_Plan_System SHALL display each Template with its name, description, and type (fixed, carb_cycle, or carb_taper)
3. THE Diet_Plan_System SHALL visually distinguish Public_Template entries from Private_Template entries using a scope label
4. WHEN a user creates a Private_Template, THE Diet_Plan_System SHALL store the Template with the creator's user_id and scope set to "private"
5. WHEN a user attempts to edit a Private_Template, THE Diet_Plan_System SHALL allow modification of the Template name, description, type, and Macro_Targets
6. WHEN a user attempts to delete a Private_Template, THE Diet_Plan_System SHALL remove the Template after user confirmation
7. IF a user attempts to edit or delete a Public_Template, THEN THE Diet_Plan_System SHALL prevent the operation and display a message indicating public templates are read-only
8. THE Diet_Plan_System SHALL provide five built-in Public_Template entries: 均衡饮食 (Balanced, fixed), 碳循环 (Carb Cycling, carb_cycle), 碳水渐降 (Carb Tapering, carb_taper), 高蛋白增肌 (High Protein, fixed), 低碳减脂 (Low Carb, fixed)

### Requirement 2: Plan Creation from Template

**User Story:** As a user, I want to create a diet plan from a template, so that I can quickly set up nutrition targets based on proven strategies.

#### Acceptance Criteria

1. WHEN a user selects a Public_Template of type "fixed", THE Diet_Plan_System SHALL auto-calculate Macro_Targets based on the user's Body_Weight and the template's predefined ratios
2. WHEN a user selects the 高蛋白增肌 template, THE Diet_Plan_System SHALL calculate protein target as 2.0 grams per kilogram of Body_Weight
3. WHEN a user selects the 均衡饮食 template, THE Diet_Plan_System SHALL calculate Macro_Targets using approximately 40% carbohydrates, 30% protein, and 30% fat of total calories
4. WHEN a user selects the 低碳减脂 template, THE Diet_Plan_System SHALL calculate Macro_Targets with a reduced carbohydrate ratio for fat loss
5. WHEN a user creates a User_Diet_Plan from a Template, THE Diet_Plan_System SHALL copy all parameters (Macro_Targets, Cycle_Config, Taper_Config) into the User_Diet_Plan as an independent snapshot
6. WHEN a Template is modified after a User_Diet_Plan has been created from the Template, THE Diet_Plan_System SHALL preserve the User_Diet_Plan parameters unchanged
7. IF the user's Body_Weight is not set in the user profile, THEN THE Diet_Plan_System SHALL prompt the user to enter Body_Weight before auto-calculating Macro_Targets
8. THE Diet_Plan_System SHALL allow the user to manually adjust auto-calculated Macro_Targets before confirming plan creation

### Requirement 3: Custom Plan Creation

**User Story:** As a user, I want to create a fully custom diet plan with manual macro targets, so that I can define my own nutrition strategy.

#### Acceptance Criteria

1. WHEN a user chooses to create a custom plan, THE Plan_Edit_Page SHALL display input fields for plan name, daily calorie target, protein target (grams), carbohydrate target (grams), and fat target (grams)
2. THE Plan_Edit_Page SHALL validate that calorie target is a positive number greater than zero
3. THE Plan_Edit_Page SHALL validate that protein, carbohydrate, and fat targets are non-negative numbers
4. WHEN the user submits a valid custom plan, THE Diet_Plan_System SHALL create a User_Diet_Plan with type "fixed" and the specified Macro_Targets
5. IF the user submits a plan with invalid or missing required fields, THEN THE Plan_Edit_Page SHALL display specific validation error messages for each invalid field

### Requirement 4: Active Plan Management

**User Story:** As a user, I want to activate and deactivate diet plans, so that only one plan guides my daily eating at a time.

#### Acceptance Criteria

1. THE Diet_Plan_System SHALL enforce that at most one User_Diet_Plan is active per user at any given time
2. WHEN a user activates a User_Diet_Plan, THE Diet_Plan_System SHALL deactivate any previously active User_Diet_Plan for that user
3. WHEN a user deactivates the current User_Diet_Plan, THE Diet_Plan_System SHALL set the plan's is_active field to false
4. THE Plan_List_Page SHALL visually highlight the currently active User_Diet_Plan
5. WHEN no User_Diet_Plan is active, THE Diet_Plan_System SHALL indicate the absence of an active plan on the Plan_List_Page

### Requirement 5: Carb Cycling Plan Configuration

**User Story:** As a user, I want to configure a carb cycling plan, so that I can alternate between high-carb training days and low-carb rest days.

#### Acceptance Criteria

1. WHEN a user selects the 碳循环 template, THE Plan_Edit_Page SHALL display a weekly calendar allowing the user to designate each day of the week as either a high-carb day or a low-carb day
2. THE Plan_Edit_Page SHALL display separate Macro_Targets for high-carb days and low-carb days
3. WHEN the Diet_Plan_System determines the current day's targets for a Carb_Cycle_Plan, THE Diet_Plan_System SHALL use the Cycle_Config to look up the current day of the week and return the corresponding Macro_Targets
4. THE Diet_Plan_System SHALL store the Cycle_Config as a mapping of weekday (0=Sunday through 6=Saturday) to day type (high or low)
5. THE Plan_Edit_Page SHALL require at least one high-carb day and at least one low-carb day in the Cycle_Config

### Requirement 6: Carb Tapering Plan Configuration

**User Story:** As a user, I want to configure a carb tapering plan, so that I can progressively reduce carbohydrate intake over multiple weeks.

#### Acceptance Criteria

1. WHEN a user selects the 碳水渐降 template, THE Plan_Edit_Page SHALL display configuration fields for total number of weeks and weekly carbohydrate reduction amount (grams)
2. THE Diet_Plan_System SHALL calculate the current week number as the number of complete weeks elapsed since the User_Diet_Plan started_at date plus one
3. THE Diet_Plan_System SHALL calculate the current week's carbohydrate target as: initial carbohydrate target minus (current week number minus one) multiplied by the weekly reduction amount
4. IF the calculated carbohydrate target falls below a minimum threshold of 50 grams, THEN THE Diet_Plan_System SHALL cap the carbohydrate target at 50 grams
5. THE Plan_Edit_Page SHALL validate that total weeks is a positive integer and weekly reduction amount is a positive number

### Requirement 7: Homepage Diet Module Display

**User Story:** As a user, I want to see my daily nutrition progress on the homepage, so that I can quickly understand how much I can still eat today.

#### Acceptance Criteria

1. WHILE a User_Diet_Plan is active, THE Homepage_Diet_Module SHALL display the Remaining_Calories value calculated as the active plan's daily calorie target minus the sum of today's recorded diet calories
2. WHILE a User_Diet_Plan is active, THE Homepage_Diet_Module SHALL display today's consumed calories and a breakdown of protein, carbohydrate, and fat intake in grams
3. WHILE a User_Diet_Plan is active, THE Homepage_Diet_Module SHALL display quick-entry buttons for each meal type: 早餐 (breakfast), 午餐 (lunch), 晚餐 (dinner), 加餐 (snack)
4. WHILE a User_Diet_Plan is active, THE Homepage_Diet_Module SHALL display a link to the Plan_List_Page for plan settings
5. WHEN no User_Diet_Plan is active, THE Homepage_Diet_Module SHALL display guidance prompting the user to set up a diet plan, with a link to the Plan_List_Page
6. WHILE a Carb_Cycle_Plan is active, THE Homepage_Diet_Module SHALL display the Macro_Targets corresponding to the current day of the week based on the Cycle_Config
7. WHILE a Carb_Taper_Plan is active, THE Homepage_Diet_Module SHALL display the Macro_Targets adjusted for the current week based on the Taper_Config and started_at date
8. IF the Remaining_Calories value is negative (consumed exceeds target), THEN THE Homepage_Diet_Module SHALL display the overage amount with a visual warning indicator

### Requirement 8: Plan Data Persistence

**User Story:** As a user, I want my diet plan data to persist across sessions, so that I do not lose my plan configuration.

#### Acceptance Criteria

1. THE Diet_Plan_System SHALL store User_Diet_Plan data in localStorage for the initial implementation phase
2. THE Diet_Plan_System SHALL store Private_Template data in localStorage for the initial implementation phase
3. WHEN the application loads, THE Diet_Plan_System SHALL restore the active User_Diet_Plan and all Private_Template entries from localStorage
4. THE Diet_Plan_System SHALL define a data structure for User_Diet_Plan containing: id, user_id, template_id (optional), is_active, name, type (fixed, carb_cycle, or carb_taper), Macro_Targets, Cycle_Config (optional), Taper_Config (optional), and started_at
5. THE Diet_Plan_System SHALL define a data structure for Template containing: id, user_id (null for public), scope (public or private), name, description, type (fixed, carb_cycle, or carb_taper), Macro_Targets, Cycle_Config (optional), and Taper_Config (optional)

### Requirement 9: Plan Editing and Deletion

**User Story:** As a user, I want to edit or delete my existing diet plans, so that I can adjust my nutrition strategy as my goals change.

#### Acceptance Criteria

1. WHEN a user edits an active User_Diet_Plan, THE Diet_Plan_System SHALL update the plan's Macro_Targets and configuration in place without creating a new plan
2. WHEN a user deletes a User_Diet_Plan, THE Diet_Plan_System SHALL remove the plan after user confirmation
3. IF a user deletes the currently active User_Diet_Plan, THEN THE Diet_Plan_System SHALL set the user's active plan state to none
4. THE Plan_List_Page SHALL provide edit and delete actions for each User_Diet_Plan owned by the current user

### Requirement 10: Page Navigation

**User Story:** As a user, I want to navigate between diet plan pages seamlessly, so that I can manage my plans without confusion.

#### Acceptance Criteria

1. THE Diet_Plan_System SHALL register the Plan_List_Page and Plan_Edit_Page as navigable pages in the application's page configuration (pages.json)
2. WHEN a user taps the plan settings link on the Homepage_Diet_Module, THE Diet_Plan_System SHALL navigate to the Plan_List_Page
3. WHEN a user taps a create or edit action on the Plan_List_Page, THE Diet_Plan_System SHALL navigate to the Plan_Edit_Page with the appropriate context (new plan or existing plan id)
4. WHEN a user completes plan creation or editing on the Plan_Edit_Page, THE Diet_Plan_System SHALL navigate back to the Plan_List_Page and refresh the plan list
