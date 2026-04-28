import { describe, it, expect, vi } from 'vitest'
import { createPinia, setActivePinia } from 'pinia'
import fc from 'fast-check'
import { PUBLIC_TEMPLATES } from '../../utils/dietPlanTemplates.js'

vi.mock('../index.js', () => ({
  useUserStore: () => ({
    userInfo: { id: 'test-user-123' },
    isLoggedIn: true
  })
}))

import { useDietPlanStore } from '../dietPlan.js'

describe('debug', () => {
  it('works', () => {
    setActivePinia(createPinia())
    const store = useDietPlanStore()
    expect(PUBLIC_TEMPLATES.length).toBe(5)
    fc.assert(fc.property(fc.integer({ min: 1, max: 100 }), (n) => {
      expect(n).toBeGreaterThan(0)
    }), { numRuns: 5 })
  })
})
