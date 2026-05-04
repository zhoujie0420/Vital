---
inclusion: auto
---

# Impeccable Design Skill

> Source: [impeccable](https://github.com/pbakaus/impeccable) by Paul Bakaus — adapted for Kiro

## Core Principle
Create distinctive, production-grade frontend interfaces with high design quality. Avoid generic "AI slop" aesthetics. Every UI should feel intentional, polished, and craft-quality.

## 7 Design Dimensions

### 1. Typography
* Build a proper type system with modular scales (not random sizes).
* Pair fonts intentionally — one display, one body maximum.
* Use OpenType features when available (tabular numbers, ligatures).
* Avoid overused fonts (Arial, default Inter, system defaults) for premium UIs.
* Control hierarchy through weight, size, AND color — not just size alone.

### 2. Color & Contrast
* Use OKLCH or HSL for perceptually uniform color manipulation.
* Tint your neutrals — pure gray looks lifeless. Add warmth or coolness.
* Dark mode is not just "invert colors" — redesign surfaces and elevations.
* Ensure WCAG AA contrast ratios minimum. Aim for AAA on body text.
* Never use gray text on colored backgrounds.

### 3. Spatial Design
* Establish a spacing system (4px or 8px base unit).
* Use consistent spacing tokens, not arbitrary pixel values.
* Visual hierarchy comes from spacing as much as from typography.
* Group related elements with tighter spacing; separate sections with generous space.

### 4. Motion Design
* Every animation needs a purpose — entrance, feedback, or state change.
* Use appropriate easing curves: ease-out for entrances, ease-in for exits.
* Stagger animations for lists and groups — never mount everything at once.
* Respect prefers-reduced-motion. Provide meaningful fallbacks.
* Avoid bounce/elastic easing — it feels dated.

### 5. Interaction Design
* Design all form states: default, focus, filled, error, disabled, loading.
* Focus states must be visible and styled (not just browser defaults).
* Loading patterns should match the layout they replace (skeleton screens).
* Provide immediate feedback for every user action.

### 6. Responsive Design
* Mobile-first approach. Design for the smallest screen, then enhance.
* Use fluid design with clamp() for typography and spacing where appropriate.
* Container queries for component-level responsiveness when available.
* Test touch targets — minimum 44x44px for interactive elements.

### 7. UX Writing
* Button labels should describe the action ("Save changes" not "Submit").
* Error messages should explain what went wrong AND how to fix it.
* Empty states should guide users toward their first action.
* Be concise. Every word should earn its place.

## Anti-Patterns to Avoid
* Don't wrap everything in cards or nest cards inside cards
* Don't use pure black/gray — always tint your neutrals
* Don't use the same border-radius everywhere
* Don't center everything by default
* Don't use placeholder text as the final copy
* Don't ignore empty, loading, and error states
* Don't use generic stock imagery
