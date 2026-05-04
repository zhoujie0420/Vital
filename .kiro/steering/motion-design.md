---
inclusion: auto
---

# Motion Design Skill

> Inspired by [Motion AI Kit](https://motion.dev/docs/ai-kit) — adapted for Kiro + uni-app/Vue

## Core Principles
Add purposeful, performant animations that enhance user experience. Every animation should serve a function: guide attention, provide feedback, or communicate state changes.

## Animation Best Practices

### Performance Tiers
Classify animations by render pipeline cost:
* **S-tier (Compositor only):** transform, opacity — use these by default
* **A-tier (Paint):** color, background-color, box-shadow — use sparingly
* **B-tier (Layout + Paint):** padding, margin, border-width — avoid animating
* **F-tier (Layout thrashing):** width, height, top, left — NEVER animate these

### CSS Transitions
* Default easing: `cubic-bezier(0.16, 1, 0.3, 1)` for smooth, natural feel
* Duration: 150-300ms for micro-interactions, 300-500ms for page transitions
* Use `will-change: transform` sparingly and only on elements that will animate

### Spring Physics
* For interactive elements, prefer spring-based animations over linear/cubic-bezier
* Natural spring: stiffness 100, damping 15-20
* Bouncy spring: stiffness 200, damping 10
* Stiff spring: stiffness 300, damping 30

### Entrance Animations
* Fade + slight translate (8-16px) for content appearing
* Stagger children with 50-100ms delay between items
* Scale from 0.95 to 1.0 for modals and overlays

### Exit Animations
* Faster than entrances (about 60-75% of entrance duration)
* Fade out with slight scale down (1.0 to 0.95)
* Use ease-in for exits (opposite of entrances)

### Scroll Animations
* Use Intersection Observer for scroll-triggered animations
* Animate elements as they enter viewport, not before
* Keep scroll animations subtle — avoid hijacking scroll behavior

### Loading States
* Skeleton screens with shimmer effect (gradient moving left to right)
* Pulse animation for placeholder content
* Progress indicators for operations > 1 second

### Micro-interactions
* Button press: scale(0.97) on active state
* Toggle switches: smooth slide with color transition
* Input focus: subtle border color transition + shadow expansion
* Success feedback: brief scale up (1.05) then back to 1.0

## uni-app / Vue Specific
* Use Vue's `<transition>` and `<transition-group>` components
* For list animations, always use `:key` with unique identifiers
* Use CSS animations over JS animations when possible for better performance
* For complex sequences, use CSS `@keyframes` with `animation-delay`

## Accessibility
* ALWAYS respect `prefers-reduced-motion: reduce`
* Provide `@media (prefers-reduced-motion: reduce)` overrides
* Ensure animations don't cause vestibular issues (no large-scale parallax)
* Keep flashing content below 3 flashes per second
