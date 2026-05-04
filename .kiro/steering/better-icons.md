---
inclusion: auto
---

# Better Icons Skill

> Inspired by [better-icons](https://github.com/better-auth/better-icons) — adapted for Kiro

## Core Principle
Use high-quality, consistent icons throughout the UI. Icons should be semantically meaningful, visually consistent, and properly sized.

## Icon Selection Rules

### Semantic Accuracy
* Choose icons that clearly represent their function
* Prefer universally recognized symbols (magnifying glass for search, gear for settings)
* When in doubt, pair icon with a text label for clarity

### Visual Consistency
* Use ONE icon library per project — do not mix icon sets
* Standardize stroke width across all icons (1.5px or 2px, pick one)
* Standardize icon size per context:
  * Navigation/Tab bar: 24-28px
  * Inline with text: match text line-height (16-20px)
  * Feature/Hero icons: 32-48px
  * Action buttons: 20-24px

### Recommended Icon Libraries
For uni-app / Vue projects:
* **uni-icons** — built-in uni-app icon component
* **Iconify** — 200+ icon sets, tree-shakeable, Vue component available
* **Custom SVG sprites** — best performance for small icon sets

For web projects:
* **Lucide** — clean, consistent, actively maintained (fork of Feather)
* **Phosphor** — flexible weight system, great for varied contexts
* **Heroicons** — by Tailwind team, pairs well with Tailwind CSS
* **Radix Icons** — minimal, designed for UI components

### Icon Usage Anti-Patterns
* NO emoji as icons — use proper SVG icons
* NO mixing filled and outlined styles in the same context
* NO icons without accessible labels (use aria-label or sr-only text)
* NO decorative icons that add visual noise without meaning
* NO icon-only buttons without tooltips (except universally understood ones like X/close)
* NO tiny icons (< 16px) — they become unreadable

### Color & State
* Default: use current text color (`currentColor`) for icons
* Interactive states: match the parent element's hover/active/focus styles
* Disabled: reduce opacity to 0.4-0.5
* Active/Selected: use accent color consistently

### Performance
* Prefer inline SVG or SVG sprites over icon fonts
* Tree-shake unused icons — don't import entire icon libraries
* For uni-app: use the built-in `<uni-icons>` component when possible
* Cache icon sprites for repeated use
