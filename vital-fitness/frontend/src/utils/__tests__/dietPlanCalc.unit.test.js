import { describe, it, expect } from 'vitest'
import { calculateMacrosFromTemplate } from '../dietPlanCalc.js'
import { PUBLIC_TEMPLATES } from '../dietPlanTemplates.js'

/**
 * Unit tests for calculateMacrosFromTemplate
 * Validates: Requirements 2.1, 2.2, 2.3, 2.4
 */

const findTemplate = (id) => PUBLIC_TEMPLATES.find(t => t.id === id)

describe('calculateMacrosFromTemplate', () => {
  describe('ratio-based template (均衡饮食)', () => {
    const template = findTemplate('tpl_balanced')

    it('calculates correct macros for bodyWeight=70kg', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      // calories = 70 * 30 = 2100
      expect(result.calories).toBe(2100)
      // protein = 2100 * 0.30 / 4 = 157.5 → 158
      expect(result.protein).toBe(158)
      // carbs = 2100 * 0.40 / 4 = 210
      expect(result.carbs).toBe(210)
      // fat = 2100 * 0.30 / 9 = 70
      expect(result.fat).toBe(70)
    })

    it('returns all integer values', () => {
      const result = calculateMacrosFromTemplate(template, 73)
      expect(Number.isInteger(result.calories)).toBe(true)
      expect(Number.isInteger(result.protein)).toBe(true)
      expect(Number.isInteger(result.carbs)).toBe(true)
      expect(Number.isInteger(result.fat)).toBe(true)
    })
  })

  describe('per-kg template (高蛋白增肌)', () => {
    const template = findTemplate('tpl_high_protein')

    it('calculates protein as 2.0g/kg body weight', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      // protein = 70 * 2.0 = 140
      expect(result.protein).toBe(140)
    })

    it('calculates correct macros for bodyWeight=70kg', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      // calories = 70 * 35 = 2450
      expect(result.calories).toBe(2450)
      // protein = 70 * 2.0 = 140
      expect(result.protein).toBe(140)
      // carbs = 2450 * 0.45 / 4 = 275.625 → 276
      expect(result.carbs).toBe(276)
      // fat = (2450 - 140*4 - 276*4) / 9 = (2450 - 560 - 1104) / 9 = 786/9 = 87.33 → 87
      expect(result.fat).toBe(87)
    })

    it('derives fat from remaining calories when fatRatio is null', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      const proteinCal = result.protein * 4
      const carbsCal = result.carbs * 4
      const remaining = result.calories - proteinCal - carbsCal
      // fat should be approximately remaining / 9
      expect(result.fat).toBe(Math.round(remaining / 9))
    })
  })

  describe('per-kg template (低碳减脂)', () => {
    const template = findTemplate('tpl_low_carb')

    it('calculates correct macros for bodyWeight=70kg', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      // calories = 70 * 24 = 1680
      expect(result.calories).toBe(1680)
      // protein = 70 * 2.0 = 140
      expect(result.protein).toBe(140)
      // carbs = 1680 * 0.20 / 4 = 84
      expect(result.carbs).toBe(84)
      // fat = (1680 - 140*4 - 84*4) / 9 = (1680 - 560 - 336) / 9 = 784/9 = 87.11 → 87
      expect(result.fat).toBe(87)
    })

    it('produces reduced carb ratio compared to balanced template', () => {
      const balanced = calculateMacrosFromTemplate(findTemplate('tpl_balanced'), 70)
      const lowCarb = calculateMacrosFromTemplate(template, 70)
      expect(lowCarb.carbs).toBeLessThan(balanced.carbs)
    })
  })

  describe('carb_cycle template (碳循环)', () => {
    const template = findTemplate('tpl_carb_cycle')

    it('returns high and low day targets', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      expect(result).toHaveProperty('high')
      expect(result).toHaveProperty('low')
      expect(result.high).toHaveProperty('calories')
      expect(result.low).toHaveProperty('calories')
    })

    it('calculates correct high day macros for bodyWeight=70kg', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      // high: calories = 70 * 33 = 2310
      expect(result.high.calories).toBe(2310)
      expect(result.high.protein).toBe(140) // 70 * 2.0
      expect(result.high.carbs).toBe(280)   // 70 * 4.0
      expect(result.high.fat).toBe(56)      // 70 * 0.8
    })
  })

  describe('carb_taper template (碳水渐降)', () => {
    const template = findTemplate('tpl_carb_taper')

    it('calculates correct initial macros for bodyWeight=70kg', () => {
      const result = calculateMacrosFromTemplate(template, 70)
      // calories = 70 * 28 = 1960
      expect(result.calories).toBe(1960)
      expect(result.protein).toBe(140)  // 70 * 2.0
      expect(result.carbs).toBe(210)    // 70 * 3.0
      expect(result.fat).toBe(63)       // 70 * 0.9
    })
  })

  describe('edge cases', () => {
    it('returns zeros for null template', () => {
      const result = calculateMacrosFromTemplate(null, 70)
      expect(result).toEqual({ calories: 0, protein: 0, carbs: 0, fat: 0 })
    })

    it('returns zeros for bodyWeight=0', () => {
      const template = findTemplate('tpl_balanced')
      const result = calculateMacrosFromTemplate(template, 0)
      expect(result).toEqual({ calories: 0, protein: 0, carbs: 0, fat: 0 })
    })

    it('returns zeros for negative bodyWeight', () => {
      const template = findTemplate('tpl_balanced')
      const result = calculateMacrosFromTemplate(template, -10)
      expect(result).toEqual({ calories: 0, protein: 0, carbs: 0, fat: 0 })
    })

    it('returns zeros for null bodyWeight', () => {
      const template = findTemplate('tpl_balanced')
      const result = calculateMacrosFromTemplate(template, null)
      expect(result).toEqual({ calories: 0, protein: 0, carbs: 0, fat: 0 })
    })

    it('returns all non-negative values for any valid input', () => {
      const templates = PUBLIC_TEMPLATES.filter(t => t.type === 'fixed')
      for (const template of templates) {
        const result = calculateMacrosFromTemplate(template, 50)
        expect(result.calories).toBeGreaterThanOrEqual(0)
        expect(result.protein).toBeGreaterThanOrEqual(0)
        expect(result.carbs).toBeGreaterThanOrEqual(0)
        expect(result.fat).toBeGreaterThanOrEqual(0)
      }
    })
  })
})
