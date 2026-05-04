---
inclusion: auto
---

# AI Kit — Design Workflow Skill

> Inspired by Motion AI Kit and Claude's frontend-design skill — adapted for Kiro

## Purpose
Guide AI-assisted UI development to produce polished, non-generic interfaces. This skill defines the workflow and quality standards for all frontend code generation.

## Design-First Workflow

### Step 1: Understand Before Building
* Read existing code and design patterns before generating new UI
* Match the project's existing style, colors, spacing, and component patterns
* Check package.json for available UI libraries before importing new ones

### Step 2: Shape Before Code
* Plan the layout structure (grid, flex, positioning) before writing markup
* Define the visual hierarchy: what should the user see first, second, third?
* Consider all states: default, loading, empty, error, success

### Step 3: Build with Craft
* Write semantic, accessible markup
* Use consistent spacing and sizing tokens
* Implement smooth transitions for state changes
* Test touch targets for mobile (min 44x44px / 80rpx)

### Step 4: Polish Before Shipping
* Check alignment and spacing consistency
* Verify color contrast meets accessibility standards
* Ensure loading and error states are handled
* Test responsive behavior across screen sizes

## Quality Checklist (Apply Before Every UI Output)
- [ ] Does the layout have clear visual hierarchy?
- [ ] Are colors consistent with the design system?
- [ ] Are all interactive states handled (hover, active, focus, disabled)?
- [ ] Are loading/empty/error states implemented?
- [ ] Is spacing consistent (using design tokens, not arbitrary values)?
- [ ] Are touch targets large enough for mobile?
- [ ] Is text readable (proper contrast, line-height, max-width)?
- [ ] Are animations purposeful and performant (transform/opacity only)?
- [ ] Does it respect prefers-reduced-motion?
- [ ] Is the code clean and well-structured?

## Common AI Design Mistakes to Avoid
1. **Generic layouts** — Don't default to centered hero + 3-column cards
2. **Inconsistent spacing** — Use a spacing scale, not random pixel values
3. **Missing states** — Always implement loading, empty, and error states
4. **Poor color choices** — Avoid neon gradients, pure black, oversaturated colors
5. **Ignoring mobile** — Design mobile-first, especially for uni-app
6. **Over-decorating** — Restraint is a design skill. Less is often more
7. **Broken images** — Use reliable placeholders, never broken URLs
8. **Generic copy** — Use realistic, contextual placeholder text
9. **No feedback** — Every user action should have visible feedback
10. **Accessibility gaps** — Labels, contrast, focus states, screen reader support
