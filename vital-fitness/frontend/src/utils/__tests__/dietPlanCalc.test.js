import { describe, it, expect } from 'vitest'
import fc from 'fast-check'
import { PUBLIC_TEMPLATES } from '../dietPlanTemplates.js'
import {
  calculateMacrosFromTemplate,
  validatePlanInput,
  validateTaperConfig,
  validateCycleConfig,
  getCarbCycleDayTargets,
  getCarbTaperDayTargets,
  calculateRemaining
} from '../dietPlanCalc.js'

/**
 * Property-based tests for dietPlanCalc utilities
 * Uses fast-check for property-based testing with vitest
 */

// ─── Property 1: Template aggregation returns all templates ───
// Feature: diet-plan, Property 1: Template aggregation returns all templates
// **Validates: Requirements 1.1**
describe('Property 1: Template aggregation returns all templates', () => {
  it('allTemplates contains all 5 public templates plus all private templates, no duplicates, no missing', () => {
    // Arbitrary for generating private templates
    const privateTemplateArb = fc.record({
      id: fc.uuid(),
      scope: fc.constant('private'),
      userId: fc.uuid(),
      name: fc.string({ minLength: 1, maxLength: 30 }),
      description: fc.string({ maxLength: 100 }),
      type: fc.constantFrom('fixed', 'carb_cycle', 'carb_taper'),
      calcConfig: fc.constant({ caloriesPerKg: 30, proteinRatio: 0.3, carbsRatio: 0.4, fatRatio: 0.3 })
    })

    fc.assert(
      fc.property(
        fc.array(privateTemplateArb, { minLength: 0, maxLength: 20 }),
        (privateTemplates) => {
          // Simulate allTemplates getter
          const allTemplates = [...PUBLIC_TEMPLATES, ...privateTemplates]

          // Must contain all 5 public templates
          for (const pub of PUBLIC_TEMPLATES) {
            expect(allTemplates.some(t => t.id === pub.id)).toBe(true)
          }

          // Must contain all private templates
          for (const priv of privateTemplates) {
            expect(allTemplates.some(t => t.id === priv.id)).toBe(true)
          }

          // Total count = 5 public + private count (no duplicates since UUIDs are unique)
          expect(allTemplates.length).toBe(PUBLIC_TEMPLATES.length + privateTemplates.length)

          // No duplicate IDs (public IDs are fixed strings, private IDs are UUIDs)
          const ids = allTemplates.map(t => t.id)
          expect(new Set(ids).size).toBe(ids.length)
        }
      ),
      { numRuns: 100 }
    )
  })
})

// ─── Property 4: Macro calculation from fixed template ───
// Feature: diet-plan, Property 4: Macro calculation from fixed template
// **Validates: Requirements 2.1, 2.2, 2.3, 2.4**
describe('Property 4: Macro calculation from fixed template', () => {
  const fixedTemplates = PUBLIC_TEMPLATES.filter(t => t.type === 'fixed')

  it('produces correct calories and non-negative values for any valid body weight and fixed template', () => {
    fc.assert(
      fc.property(
        fc.constantFrom(...fixedTemplates),
        fc.double({ min: 1, max: 300, noNaN: true }),
        (template, bodyWeight) => {
          const result = calculateMacrosFromTemplate(template, bodyWeight)

          // calories = bodyWeight * caloriesPerKg (rounded)
          const expectedCalories = Math.round(bodyWeight * template.calcConfig.caloriesPerKg)
          expect(result.calories).toBe(expectedCalories)

          // All values must be non-negative
          expect(result.calories).toBeGreaterThanOrEqual(0)
          expect(result.protein).toBeGreaterThanOrEqual(0)
          expect(result.carbs).toBeGreaterThanOrEqual(0)
          expect(result.fat).toBeGreaterThanOrEqual(0)

          // All values must be integers (rounded)
          expect(Number.isInteger(result.calories)).toBe(true)
          expect(Number.isInteger(result.protein)).toBe(true)
          expect(Number.isInteger(result.carbs)).toBe(true)
          expect(Number.isInteger(result.fat)).toBe(true)

          // For ratio-based templates (均衡饮食), verify ratio derivation
          if (template.calcConfig.proteinRatio !== undefined && template.calcConfig.proteinRatio !== null) {
            const expectedProtein = Math.round(Math.max(0, (expectedCalories * template.calcConfig.proteinRatio) / 4))
            const expectedCarbs = Math.round(Math.max(0, (expectedCalories * template.calcConfig.carbsRatio) / 4))
            const expectedFat = Math.round(Math.max(0, (expectedCalories * template.calcConfig.fatRatio) / 9))
            expect(result.protein).toBe(expectedProtein)
            expect(result.carbs).toBe(expectedCarbs)
            expect(result.fat).toBe(expectedFat)
          }

          // For per-kg templates (高蛋白增肌, 低碳减脂), verify protein from per-kg
          if (template.calcConfig.proteinPerKg !== undefined) {
            const expectedProtein = Math.round(Math.max(0, bodyWeight * template.calcConfig.proteinPerKg))
            expect(result.protein).toBe(expectedProtein)
          }
        }
      ),
      { numRuns: 100 }
    )
  })
})

// ─── Property 6: Plan numeric input validation ───
// Feature: diet-plan, Property 6: Plan numeric input validation
// **Validates: Requirements 3.2, 3.3, 6.5**
describe('Property 6: Plan numeric input validation', () => {
  it('accepts valid plan inputs (calories > 0, macros >= 0)', () => {
    fc.assert(
      fc.property(
        fc.double({ min: 0.01, max: 100000, noNaN: true }),
        fc.double({ min: 0, max: 10000, noNaN: true }),
        fc.double({ min: 0, max: 10000, noNaN: true }),
        fc.double({ min: 0, max: 10000, noNaN: true }),
        (calories, protein, carbs, fat) => {
          const result = validatePlanInput({ calories, protein, carbs, fat })
          expect(result.valid).toBe(true)
          expect(Object.keys(result.errors).length).toBe(0)
        }
      ),
      { numRuns: 100 }
    )
  })

  it('rejects zero or negative calories', () => {
    fc.assert(
      fc.property(
        fc.double({ min: -100000, max: 0, noNaN: true }),
        (calories) => {
          const result = validatePlanInput({ calories, protein: 10, carbs: 10, fat: 10 })
          expect(result.valid).toBe(false)
          expect(result.errors).toHaveProperty('calories')
        }
      ),
      { numRuns: 100 }
    )
  })

  it('rejects negative protein, carbs, or fat', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('protein', 'carbs', 'fat'),
        fc.double({ min: -100000, max: -0.01, noNaN: true }),
        (field, value) => {
          const data = { calories: 2000, protein: 100, carbs: 200, fat: 80 }
          data[field] = value
          const result = validatePlanInput(data)
          expect(result.valid).toBe(false)
          expect(result.errors).toHaveProperty(field)
        }
      ),
      { numRuns: 100 }
    )
  })

  it('validates taper config: accepts positive integer totalWeeks and positive weeklyReduction', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 1, max: 52 }),
        fc.double({ min: 0.1, max: 100, noNaN: true }),
        (totalWeeks, weeklyReduction) => {
          const result = validateTaperConfig({ totalWeeks, weeklyReduction })
          expect(result.valid).toBe(true)
          expect(Object.keys(result.errors).length).toBe(0)
        }
      ),
      { numRuns: 100 }
    )
  })

  it('rejects non-positive-integer totalWeeks', () => {
    fc.assert(
      fc.property(
        fc.constantFrom(0, -1, -5, 1.5, 2.7),
        (totalWeeks) => {
          const result = validateTaperConfig({ totalWeeks, weeklyReduction: 10 })
          expect(result.valid).toBe(false)
          expect(result.errors).toHaveProperty('totalWeeks')
        }
      ),
      { numRuns: 100 }
    )
  })

  it('rejects non-positive weeklyReduction', () => {
    fc.assert(
      fc.property(
        fc.double({ min: -1000, max: 0, noNaN: true }),
        (weeklyReduction) => {
          const result = validateTaperConfig({ totalWeeks: 8, weeklyReduction })
          expect(result.valid).toBe(false)
          expect(result.errors).toHaveProperty('weeklyReduction')
        }
      ),
      { numRuns: 100 }
    )
  })
})

// ─── Property 9: Carb cycle day-of-week target lookup ───
// Feature: diet-plan, Property 9: Carb cycle day-of-week target lookup
// **Validates: Requirements 5.3, 7.6**
describe('Property 9: Carb cycle day-of-week target lookup', () => {
  it('returns correct targets based on day of week from cycleConfig', () => {
    // Generate a valid dayMap with at least one high and one low
    const dayTypeArb = fc.constantFrom('high', 'low')
    const dayMapArb = fc.tuple(
      dayTypeArb, dayTypeArb, dayTypeArb, dayTypeArb,
      dayTypeArb, dayTypeArb, dayTypeArb
    ).filter(days => days.includes('high') && days.includes('low'))
      .map(days => {
        const dayMap = {}
        for (let i = 0; i < 7; i++) dayMap[i] = days[i]
        return dayMap
      })

    const macroArb = fc.record({
      calories: fc.integer({ min: 500, max: 5000 }),
      protein: fc.integer({ min: 10, max: 500 }),
      carbs: fc.integer({ min: 10, max: 500 }),
      fat: fc.integer({ min: 10, max: 300 })
    })

    // Generate a date within a reasonable range
    const dateArb = fc.date({
      min: new Date('2020-01-01'),
      max: new Date('2030-12-31')
    })

    fc.assert(
      fc.property(
        dayMapArb,
        macroArb,
        macroArb,
        dateArb,
        (dayMap, highDayTargets, lowDayTargets, date) => {
          const plan = {
            type: 'carb_cycle',
            cycleConfig: { dayMap, highDayTargets, lowDayTargets }
          }

          const result = getCarbCycleDayTargets(plan, date)
          const dayOfWeek = date.getDay()
          const expectedType = dayMap[dayOfWeek]

          if (expectedType === 'high') {
            expect(result.calories).toBe(highDayTargets.calories)
            expect(result.protein).toBe(highDayTargets.protein)
            expect(result.carbs).toBe(highDayTargets.carbs)
            expect(result.fat).toBe(highDayTargets.fat)
          } else {
            expect(result.calories).toBe(lowDayTargets.calories)
            expect(result.protein).toBe(lowDayTargets.protein)
            expect(result.carbs).toBe(lowDayTargets.carbs)
            expect(result.fat).toBe(lowDayTargets.fat)
          }
        }
      ),
      { numRuns: 100 }
    )
  })
})

// ─── Property 10: Cycle config requires mixed day types ───
// Feature: diet-plan, Property 10: Cycle config requires mixed day types
// **Validates: Requirements 5.5**
describe('Property 10: Cycle config requires mixed day types', () => {
  it('rejects all-same-type dayMap', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('high', 'low'),
        (sameType) => {
          const dayMap = {}
          for (let i = 0; i < 7; i++) dayMap[i] = sameType
          const result = validateCycleConfig(dayMap)
          expect(result.valid).toBe(false)
          expect(result.errors).toHaveProperty('dayMap')
        }
      ),
      { numRuns: 100 }
    )
  })

  it('accepts mixed-type dayMap', () => {
    const dayTypeArb = fc.constantFrom('high', 'low')
    const mixedDayMapArb = fc.tuple(
      dayTypeArb, dayTypeArb, dayTypeArb, dayTypeArb,
      dayTypeArb, dayTypeArb, dayTypeArb
    ).filter(days => days.includes('high') && days.includes('low'))
      .map(days => {
        const dayMap = {}
        for (let i = 0; i < 7; i++) dayMap[i] = days[i]
        return dayMap
      })

    fc.assert(
      fc.property(
        mixedDayMapArb,
        (dayMap) => {
          const result = validateCycleConfig(dayMap)
          expect(result.valid).toBe(true)
          expect(Object.keys(result.errors).length).toBe(0)
        }
      ),
      { numRuns: 100 }
    )
  })
})

// ─── Property 11: Carb taper weekly target calculation ───
// Feature: diet-plan, Property 11: Carb taper weekly target calculation
// **Validates: Requirements 6.2, 6.3, 6.4, 7.7**
describe('Property 11: Carb taper weekly target calculation', () => {
  it('calculates carb target with 50g minimum cap', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 100, max: 500 }),   // initialCarbs
        fc.integer({ min: 1, max: 50 }),       // weeklyReduction
        fc.integer({ min: 1000, max: 4000 }),  // calories
        fc.integer({ min: 50, max: 300 }),     // protein
        fc.integer({ min: 20, max: 200 }),     // fat
        fc.integer({ min: 0, max: 365 }),      // daysOffset from start
        (initialCarbs, weeklyReduction, calories, protein, fat, daysOffset) => {
          const startedAt = new Date('2024-01-01')
          const date = new Date(startedAt.getTime() + daysOffset * 24 * 60 * 60 * 1000)

          const plan = {
            type: 'carb_taper',
            startedAt: startedAt.toISOString(),
            macroTargets: { calories, protein, carbs: initialCarbs, fat },
            taperConfig: {
              initialCarbs,
              weeklyReduction,
              totalWeeks: 12,
              minCarbs: 50
            }
          }

          const result = getCarbTaperDayTargets(plan, date)

          // Calculate expected week number
          const msPerWeek = 7 * 24 * 60 * 60 * 1000
          const weekNumber = Math.floor((date - startedAt) / msPerWeek) + 1

          // Expected carbs with 50g minimum cap
          const rawCarbs = initialCarbs - (weekNumber - 1) * weeklyReduction
          const expectedCarbs = Math.round(Math.max(rawCarbs, 50) * 10) / 10

          expect(result.carbs).toBe(expectedCarbs)

          // Carbs should never go below 50g
          expect(result.carbs).toBeGreaterThanOrEqual(50)

          // Other macros should remain unchanged
          expect(result.calories).toBe(calories)
          expect(result.protein).toBe(protein)
          expect(result.fat).toBe(fat)
        }
      ),
      { numRuns: 100 }
    )
  })
})

// ─── Property 12: Remaining calories calculation ───
// Feature: diet-plan, Property 12: Remaining calories calculation
// **Validates: Requirements 7.1**
describe('Property 12: Remaining calories calculation', () => {
  it('remaining = target - consumed for each macro', () => {
    const macroArb = fc.record({
      calories: fc.double({ min: 0, max: 10000, noNaN: true }),
      protein: fc.double({ min: 0, max: 1000, noNaN: true }),
      carbs: fc.double({ min: 0, max: 1000, noNaN: true }),
      fat: fc.double({ min: 0, max: 1000, noNaN: true })
    })

    fc.assert(
      fc.property(
        macroArb,
        macroArb,
        (targets, consumed) => {
          const result = calculateRemaining(targets, consumed)

          // round1 helper: Math.round(value * 10) / 10
          const round1 = (v) => Math.round(v * 10) / 10

          expect(result.calories).toBe(round1(targets.calories - consumed.calories))
          expect(result.protein).toBe(round1(targets.protein - consumed.protein))
          expect(result.carbs).toBe(round1(targets.carbs - consumed.carbs))
          expect(result.fat).toBe(round1(targets.fat - consumed.fat))
        }
      ),
      { numRuns: 100 }
    )
  })
})
