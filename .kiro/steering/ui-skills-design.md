---
inclusion: auto
---

# UI Skills Design System

> Inspired by Claude's DESIGN.md concept and [awesome-claude-design](https://github.com/VoltAgent/awesome-claude-design) — adapted for Kiro + uni-app/Vue fitness app

## Design System Tokens

### Colors
```
Primary:       #4CAF50 (Health Green)
Primary Dark:  #388E3C
Primary Light: #C8E6C9
Accent:        #FF6B35 (Energy Orange)
Background:    #FAFAFA
Surface:       #FFFFFF
Text Primary:  #1A1A2E
Text Secondary:#6B7280
Text Muted:    #9CA3AF
Border:        #E5E7EB
Success:       #10B981
Warning:       #F59E0B
Error:         #EF4444
```

### Spacing Scale (8px base)
```
xs:   4px
sm:   8px
md:   16px
lg:   24px
xl:   32px
2xl:  48px
3xl:  64px
```

### Border Radius
```
sm:   8rpx
md:   16rpx
lg:   24rpx
xl:   32rpx
full: 9999rpx (pills, avatars)
```

### Typography Scale (uni-app rpx)
```
Display:  48rpx  / font-weight: 700 / line-height: 1.1
H1:       40rpx  / font-weight: 700 / line-height: 1.2
H2:       34rpx  / font-weight: 600 / line-height: 1.3
H3:       30rpx  / font-weight: 600 / line-height: 1.4
Body:     28rpx  / font-weight: 400 / line-height: 1.6
Caption:  24rpx  / font-weight: 400 / line-height: 1.5
Small:    22rpx  / font-weight: 400 / line-height: 1.4
```

### Shadows
```
sm:   0 2rpx 8rpx rgba(0, 0, 0, 0.06)
md:   0 4rpx 16rpx rgba(0, 0, 0, 0.08)
lg:   0 8rpx 32rpx rgba(0, 0, 0, 0.10)
```

## Component Patterns

### Cards
* Use rounded corners (16rpx default)
* Subtle shadow (sm) for elevation
* Consistent internal padding (24rpx-32rpx)
* Clear visual hierarchy: title → subtitle → content → action

### Buttons
* Primary: filled background, white text, rounded
* Secondary: outlined or ghost style
* Min height: 80rpx for touch targets
* Include loading state with spinner
* Disabled state: reduced opacity (0.5)

### Form Inputs
* Label above input (always)
* Clear focus state with accent border
* Error state: red border + error message below
* Consistent height: 80-88rpx
* Placeholder text in muted color

### Lists
* Consistent item height and padding
* Dividers between items (1rpx border or spacing)
* Swipe actions where appropriate
* Pull-to-refresh for data lists

### Navigation
* Tab bar: 4-5 items max, icon + label
* Active state: accent color, slightly larger icon
* Page headers: left-aligned title, right action buttons

## Fitness App Specific Patterns

### Data Visualization
* Use progress rings/bars for goals
* Color-code metrics (green=good, yellow=warning, red=attention)
* Show trends with simple line charts
* Display units clearly (kg, kcal, min)

### Workout Cards
* Show exercise name, sets, reps prominently
* Use icons for exercise type (strength, cardio, flexibility)
* Timer/duration in large, readable format
* Completion status with checkmarks

### Diet/Nutrition
* Macro breakdown with colored progress bars
* Meal cards grouped by time (breakfast, lunch, dinner, snack)
* Calorie summary prominent at top
* Food item list with portion sizes

### Mood Tracking
* Use expressive but tasteful mood indicators (not emoji)
* Color-coded mood scale
* Calendar/timeline view for mood history
* Simple, quick-entry interface (minimize friction)

### Weight Tracking
* Large current weight display
* Trend line chart showing progress
* Goal indicator on chart
* BMI or other derived metrics as secondary info
