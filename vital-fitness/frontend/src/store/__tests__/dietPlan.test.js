/**
 * Property-based tests and unit tests for useDietPlanStore
 */
import { describe, it, expect, beforeEach, vi } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import fc from 'fast-check'
import { PUBLIC_TEMPLATES } from '../../utils/dietPlanTemplates.js'

vi.mock('../index.js', () => ({
  useUserStore: () => ({ userInfo: { id: 'test-user-123' }, isLoggedIn: true })
}))

import { useDietPlanStore } from '../dietPlan.js'

function resetStorage() {
  const s = globalThis.__testMockStorage
  if (s) for (const k of Object.keys(s)) delete s[k]
}

function setup() {
  setActivePinia(createPinia())
  return useDietPlanStore()
}

describe('dietPlanStore', () => {
  beforeEach(resetStorage)

  // Property 2
  it('createPrivateTemplate sets scope=private and userId', () => {
    fc.assert(fc.property(
      fc.record({
        name: fc.string({ minLength: 1, maxLength: 30 }),
        description: fc.string({ maxLength: 50 }),
        type: fc.constantFrom('fixed', 'carb_cycle', 'carb_taper'),
        calcConfig: fc.constant({ caloriesPerKg: 30, proteinRatio: 0.3, carbsRatio: 0.4, fatRatio: 0.3 })
      }),
      (data) => {
        const store = setup()
        const r = store.createPrivateTemplate({ ...data })
        expect(r.scope).toBe('private')
        expect(r.userId).toBe('test-user-123')
        expect(store.privateTemplates.some(t => t.id === r.id)).toBe(true)
      }
    ), { numRuns: 50 })
  })

  // Property 3
  it('public templates cannot be updated or deleted', () => {
    for (const pub of PUBLIC_TEMPLATES) {
      const store = setup()
      expect(() => store.updatePrivateTemplate(pub.id, { name: 'x' })).toThrow()
      expect(() => store.deletePrivateTemplate(pub.id)).toThrow()
    }
  })

  // Property 5
  it('template-to-plan snapshot isolation', () => {
    const store = setup()
    const tpl = store.createPrivateTemplate({
      name: 'T', description: '', type: 'fixed',
      calcConfig: { caloriesPerKg: 30, proteinRatio: 0.3, carbsRatio: 0.4, fatRatio: 0.3 }
    })
    const plan = store.createPlanFromTemplate(tpl.id, {}, 70)
    const before = { ...plan.macroTargets }
    store.updatePrivateTemplate(tpl.id, { calcConfig: { caloriesPerKg: 99 } })
    expect(store.plans.find(p => p.id === plan.id).macroTargets).toEqual(before)
  })

  // Property 7
  it('createCustomPlan produces fixed-type plan', () => {
    fc.assert(fc.property(
      fc.record({
        name: fc.string({ minLength: 1, maxLength: 20 }),
        calories: fc.double({ min: 0.01, max: 5000, noNaN: true }),
        protein: fc.double({ min: 0, max: 500, noNaN: true }),
        carbs: fc.double({ min: 0, max: 500, noNaN: true }),
        fat: fc.double({ min: 0, max: 300, noNaN: true })
      }),
      (d) => {
        const store = setup()
        const p = store.createCustomPlan({ ...d })
        expect(p.type).toBe('fixed')
        expect(p.macroTargets.calories).toBe(d.calories)
      }
    ), { numRuns: 50 })
  })

  // Property 8
  it('single active plan invariant', () => {
    const store = setup()
    const ids = []
    for (let i = 0; i < 4; i++) {
      ids.push(store.createCustomPlan({ name: 'P' + i, calories: 2000, protein: 100, carbs: 200, fat: 80 }).id)
    }
    store.activatePlan(ids[0])
    store.activatePlan(ids[2])
    expect(store.plans.filter(p => p.isActive).length).toBe(1)
    expect(store.activePlanId).toBe(ids[2])
    store.deactivatePlan(ids[2])
    expect(store.plans.filter(p => p.isActive).length).toBe(0)
    expect(store.activePlanId).toBeNull()
  })

  // Property 13
  it('localStorage persistence round-trip', () => {
    const store = setup()
    const plan = store.createCustomPlan({ name: 'RT', calories: 2000, protein: 100, carbs: 200, fat: 80 })
    store.createPrivateTemplate({ name: 'TPL', description: '', type: 'fixed', calcConfig: {} })
    store.activatePlan(plan.id)
    const snap = JSON.parse(JSON.stringify({ plans: store.plans, tpls: store.privateTemplates, active: store.activePlanId }))
    const store2 = setup()
    store2.init()
    expect(store2.plans).toEqual(snap.plans)
    expect(store2.privateTemplates).toEqual(snap.tpls)
    expect(store2.activePlanId).toBe(snap.active)
  })

  // Property 14
  it('updatePlan preserves id, userId, createdAt', () => {
    const store = setup()
    const p = store.createCustomPlan({ name: 'Old', calories: 2000, protein: 100, carbs: 200, fat: 80 })
    const { id, userId, createdAt } = p
    store.updatePlan(p.id, { name: 'New', macroTargets: { calories: 3000, protein: 150, carbs: 300, fat: 100 } })
    const u = store.plans.find(x => x.id === id)
    expect(u.id).toBe(id)
    expect(u.userId).toBe(userId)
    expect(u.createdAt).toBe(createdAt)
    expect(u.name).toBe('New')
  })

  // Property 15
  it('deleting active plan clears active state', () => {
    const store = setup()
    const p = store.createCustomPlan({ name: 'Del', calories: 2000, protein: 100, carbs: 200, fat: 80 })
    store.activatePlan(p.id)
    store.deletePlan(p.id)
    expect(store.activePlanId).toBeNull()
    expect(store.activePlan).toBeNull()
  })

  // CRUD unit tests
  it('carb_cycle plan has cycleConfig', () => {
    const store = setup()
    const plan = store.createPlanFromTemplate('tpl_carb_cycle', {}, 70)
    expect(plan.cycleConfig.dayMap).toBeDefined()
    expect(plan.cycleConfig.highDayTargets).toBeDefined()
  })

  it('deletePlan updates localStorage', () => {
    const store = setup()
    const p1 = store.createCustomPlan({ name: 'A', calories: 2000, protein: 100, carbs: 200, fat: 80 })
    const p2 = store.createCustomPlan({ name: 'B', calories: 2500, protein: 120, carbs: 250, fat: 90 })
    store.deletePlan(p1.id)
    expect(store.plans.length).toBe(1)
    expect(store.plans[0].id).toBe(p2.id)
  })

  it('corrupted localStorage handled gracefully', () => {
    const st = globalThis.__testMockStorage
    st['diet_plans'] = '{bad'
    st['private_templates'] = 'nope'
    st['active_plan_id'] = ''
    const spy = vi.spyOn(console, 'error').mockImplementation(() => {})
    const store = setup()
    store.init()
    spy.mockRestore()
    expect(store.plans).toEqual([])
    expect(store.privateTemplates).toEqual([])
  })
})
