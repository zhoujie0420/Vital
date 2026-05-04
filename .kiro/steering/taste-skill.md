---
inclusion: auto
---

# High-Agency Frontend Design Skill (Taste)

> Source: [taste-skill](https://github.com/Leonxlnx/taste-skill) — adapted for Kiro + uni-app/Vue

## 1. ACTIVE BASELINE CONFIGURATION
* DESIGN_VARIANCE: 7 (1=Perfect Symmetry, 10=Artsy Chaos)
* MOTION_INTENSITY: 5 (1=Static/No movement, 10=Cinematic)
* VISUAL_DENSITY: 5 (1=Art Gallery/Airy, 10=Pilot Cockpit/Packed Data)

Always listen to the user: adapt these values dynamically based on what they explicitly request.

## 2. DESIGN ENGINEERING DIRECTIVES (Bias Correction)

**Rule 1: Deterministic Typography**
* Headlines: Use tight tracking and controlled leading for premium feel.
* Body: Use comfortable line-height and constrain max-width for readability.
* ANTI-SLOP: Avoid default system fonts for premium UIs. Choose distinctive, high-quality fonts.
* For technical/dashboard UIs, use exclusively Sans-Serif pairings.

**Rule 2: Color Calibration**
* Max 1 Accent Color. Saturation < 80%.
* THE PURPLE BAN: The "AI Purple/Blue" aesthetic is banned. No purple button glows, no neon gradients.
* Use neutral bases with high-contrast, singular accents (e.g. Emerald, Electric Blue, or Deep Rose).
* Stick to one palette for the entire output. Do not fluctuate between warm and cool grays.

**Rule 3: Layout Diversification**
* ANTI-CENTER BIAS: Avoid always centering everything. Use asymmetric layouts, split screens, or left-aligned content when appropriate.

**Rule 4: Materiality & Shadows**
* Use cards ONLY when elevation communicates hierarchy.
* When a shadow is used, tint it to the background hue.
* For dense UIs, prefer border-top, divide-y, or negative space over card containers.

**Rule 5: Interactive UI States**
* MUST implement full interaction cycles:
  * Loading: Skeleton loaders matching layout sizes (avoid generic circular spinners).
  * Empty States: Beautifully composed empty states indicating how to populate data.
  * Error States: Clear, inline error reporting.
  * Tactile Feedback: On press, use subtle scale or translate to simulate physical push.

**Rule 6: Data & Form Patterns**
* Label MUST sit above input. Helper text optional. Error text below input.

## 3. AI TELLS (Forbidden Patterns)

### Visual & CSS
* NO Neon/Outer Glows — use inner borders or subtle tinted shadows
* NO Pure Black (#000000) — use off-black or deep gray
* NO Oversaturated Accents — desaturate to blend with neutrals
* NO Excessive Gradient Text on large headers

### Typography
* NO Oversized H1s — control hierarchy with weight and color, not just massive scale

### Layout & Spacing
* Align & Space Perfectly — ensure padding and margins are mathematically consistent
* NO 3-Column Equal Card Layouts as default — use varied grid layouts

### Content & Data
* NO Generic Names ("John Doe") — use creative, realistic names
* NO Fake Numbers (99.99%, 50%) — use organic data (47.2%, 83.7%)
* NO Filler Words ("Elevate", "Seamless", "Unleash", "Next-Gen") — use concrete verbs

### External Resources
* NO Broken Image Links — use reliable placeholders
* Production-Ready Cleanliness: Code must be clean, visually striking, and refined

## 4. PERFORMANCE GUARDRAILS
* Apply grain/noise filters only to fixed, pointer-events-none elements
* Never animate top, left, width, or height — use transform and opacity
* Avoid arbitrary z-index spam — use z-indexes strictly for systemic layer contexts
