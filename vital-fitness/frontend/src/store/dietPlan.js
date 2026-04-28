/**
 * 饮食规划 Pinia Store
 * Diet plan state management
 *
 * 管理用户饮食方案和自定义模板，数据通过 localStorage 持久化。
 * 遵循现有 useUserStore 的 Options API 风格。
 */
import { defineStore } from 'pinia'
import { PUBLIC_TEMPLATES } from '../utils/dietPlanTemplates.js'
import { calculateMacrosFromTemplate, validatePlanInput } from '../utils/dietPlanCalc.js'
import { useUserStore } from './index.js'

/**
 * 生成简易 UUID
 * @returns {string}
 */
function generateId() {
  return Date.now().toString(36) + Math.random().toString(36).substr(2)
}

// localStorage keys
const STORAGE_KEY_PLANS = 'diet_plans'
const STORAGE_KEY_PRIVATE_TEMPLATES = 'private_templates'
const STORAGE_KEY_ACTIVE_PLAN_ID = 'active_plan_id'

export const useDietPlanStore = defineStore('dietPlan', {
  state: () => ({
    /** @type {import('../utils/dietPlanTemplates').UserDietPlan[]} */
    plans: [],
    /** @type {import('../utils/dietPlanTemplates').DietPlanTemplate[]} */
    privateTemplates: [],
    /** @type {string|null} */
    activePlanId: null
  }),

  getters: {
    /**
     * 当前激活的方案
     * @returns {object|null}
     */
    activePlan: (state) => state.plans.find(p => p.id === state.activePlanId) || null,

    /**
     * 所有模板（公共 + 私有）
     * @returns {object[]}
     */
    allTemplates: (state) => [...PUBLIC_TEMPLATES, ...state.privateTemplates],

    /**
     * 公共模板
     * @returns {object[]}
     */
    publicTemplates: () => PUBLIC_TEMPLATES
  },

  actions: {
    // ==================== 初始化 ====================

    /**
     * 初始化 store，从 localStorage 恢复数据
     */
    init() {
      this._restore()
    },

    // ==================== 持久化 ====================

    /**
     * 将当前状态持久化到 localStorage
     */
    _persist() {
      try {
        uni.setStorageSync(STORAGE_KEY_PLANS, JSON.stringify(this.plans))
        uni.setStorageSync(STORAGE_KEY_PRIVATE_TEMPLATES, JSON.stringify(this.privateTemplates))
        uni.setStorageSync(STORAGE_KEY_ACTIVE_PLAN_ID, this.activePlanId || '')
      } catch (e) {
        console.error('Failed to persist diet plan data:', e)
      }
    },

    /**
     * 从 localStorage 恢复数据，JSON.parse 失败时重置为空
     */
    _restore() {
      try {
        const plansStr = uni.getStorageSync(STORAGE_KEY_PLANS)
        this.plans = plansStr ? JSON.parse(plansStr) : []
      } catch (e) {
        console.error('Failed to restore diet plans:', e)
        this.plans = []
      }

      try {
        const templatesStr = uni.getStorageSync(STORAGE_KEY_PRIVATE_TEMPLATES)
        this.privateTemplates = templatesStr ? JSON.parse(templatesStr) : []
      } catch (e) {
        console.error('Failed to restore private templates:', e)
        this.privateTemplates = []
      }

      try {
        const activeId = uni.getStorageSync(STORAGE_KEY_ACTIVE_PLAN_ID)
        this.activePlanId = activeId || null
      } catch (e) {
        console.error('Failed to restore active plan id:', e)
        this.activePlanId = null
      }
    },

    // ==================== 方案 CRUD ====================

    /**
     * 从模板创建方案
     *
     * @param {string} templateId - 模板 ID
     * @param {object} [overrides={}] - 用户手动调整的覆盖值
     * @param {number} bodyWeight - 用户体重（公斤）
     * @returns {object} 创建的方案
     */
    createPlanFromTemplate(templateId, overrides = {}, bodyWeight) {
      // 在所有模板中查找（公共 + 私有）
      const template = this.allTemplates.find(t => t.id === templateId)
      if (!template) {
        throw new Error(`Template not found: ${templateId}`)
      }

      const userStore = useUserStore()
      const userId = userStore.userInfo?.id || 'anonymous'
      const now = new Date().toISOString()

      // 计算宏量营养素
      const calculatedMacros = calculateMacrosFromTemplate(template, bodyWeight)

      // 构建方案对象
      const plan = {
        id: generateId(),
        userId,
        templateId: template.id,
        isActive: false,
        name: overrides.name || template.name,
        type: template.type,
        startedAt: overrides.startedAt || now,
        createdAt: now,
        macroTargets: null,
        cycleConfig: null,
        taperConfig: null
      }

      // 根据类型设置目标和配置（快照复制，与模板隔离）
      if (template.type === 'carb_cycle') {
        const highTargets = overrides.highDayTargets || calculatedMacros.high || { calories: 0, protein: 0, carbs: 0, fat: 0 }
        const lowTargets = overrides.lowDayTargets || calculatedMacros.low || { calories: 0, protein: 0, carbs: 0, fat: 0 }
        plan.macroTargets = { ...highTargets } // 使用高碳日作为基础目标
        plan.cycleConfig = {
          dayMap: overrides.dayMap
            ? { ...overrides.dayMap }
            : (template.defaultCycleConfig ? { ...template.defaultCycleConfig } : {}),
          highDayTargets: { ...highTargets },
          lowDayTargets: { ...lowTargets }
        }
      } else if (template.type === 'carb_taper') {
        const macros = overrides.macroTargets
          ? { ...overrides.macroTargets }
          : { ...calculatedMacros }
        plan.macroTargets = macros
        const defaultTaper = template.defaultTaperConfig || {}
        plan.taperConfig = {
          totalWeeks: overrides.totalWeeks || defaultTaper.totalWeeks || 8,
          weeklyReduction: overrides.weeklyReduction || defaultTaper.weeklyReduction || 15,
          initialCarbs: overrides.initialCarbs || macros.carbs,
          minCarbs: overrides.minCarbs || 50
        }
      } else {
        // fixed type
        plan.macroTargets = overrides.macroTargets
          ? { ...overrides.macroTargets }
          : { ...calculatedMacros }
      }

      this.plans.push(plan)
      this._persist()
      return plan
    },

    /**
     * 创建自定义方案
     *
     * @param {object} planData - 方案数据，包含 name, calories, protein, carbs, fat
     * @returns {object} 创建的方案
     */
    createCustomPlan(planData) {
      const validation = validatePlanInput({
        calories: planData.calories,
        protein: planData.protein,
        carbs: planData.carbs,
        fat: planData.fat
      })

      if (!validation.valid) {
        throw new Error(`Invalid plan input: ${JSON.stringify(validation.errors)}`)
      }

      const userStore = useUserStore()
      const userId = userStore.userInfo?.id || 'anonymous'
      const now = new Date().toISOString()

      const plan = {
        id: generateId(),
        userId,
        templateId: null,
        isActive: false,
        name: planData.name || '自定义方案',
        type: 'fixed',
        startedAt: now,
        createdAt: now,
        macroTargets: {
          calories: planData.calories,
          protein: planData.protein,
          carbs: planData.carbs,
          fat: planData.fat
        },
        cycleConfig: null,
        taperConfig: null
      }

      this.plans.push(plan)
      this._persist()
      return plan
    },

    /**
     * 更新方案（保留 id, userId, createdAt）
     *
     * @param {string} planId - 方案 ID
     * @param {object} updates - 更新字段
     * @returns {object} 更新后的方案
     */
    updatePlan(planId, updates) {
      const index = this.plans.findIndex(p => p.id === planId)
      if (index === -1) {
        throw new Error(`Plan not found: ${planId}`)
      }

      const plan = this.plans[index]
      // 保留不可变字段
      const preserved = {
        id: plan.id,
        userId: plan.userId,
        createdAt: plan.createdAt
      }

      // 合并更新
      Object.assign(plan, updates, preserved)
      this._persist()
      return plan
    },

    /**
     * 删除方案
     *
     * @param {string} planId - 方案 ID
     */
    deletePlan(planId) {
      const index = this.plans.findIndex(p => p.id === planId)
      if (index === -1) {
        throw new Error(`Plan not found: ${planId}`)
      }

      // 如果删除的是激活方案，清除激活状态
      if (this.activePlanId === planId) {
        this.activePlanId = null
      }

      this.plans.splice(index, 1)
      this._persist()
    },

    // ==================== 激活/停用 ====================

    /**
     * 激活方案（确保单一激活不变量）
     *
     * @param {string} planId - 方案 ID
     */
    activatePlan(planId) {
      const target = this.plans.find(p => p.id === planId)
      if (!target) {
        throw new Error(`Plan not found: ${planId}`)
      }

      // 先将所有方案设为非激活
      this.plans.forEach(p => {
        p.isActive = false
      })

      // 激活目标方案
      target.isActive = true
      this.activePlanId = planId
      this._persist()
    },

    /**
     * 停用方案
     *
     * @param {string} planId - 方案 ID
     */
    deactivatePlan(planId) {
      const target = this.plans.find(p => p.id === planId)
      if (!target) {
        throw new Error(`Plan not found: ${planId}`)
      }

      target.isActive = false
      this.activePlanId = null
      this._persist()
    },

    // ==================== 自定义模板 CRUD ====================

    /**
     * 创建自定义模板
     *
     * @param {object} templateData - 模板数据
     * @returns {object} 创建的模板
     */
    createPrivateTemplate(templateData) {
      const userStore = useUserStore()
      const userId = userStore.userInfo?.id || 'anonymous'

      const template = {
        ...templateData,
        id: generateId(),
        scope: 'private',
        userId
      }

      this.privateTemplates.push(template)
      this._persist()
      return template
    },

    /**
     * 更新自定义模板（拒绝公共模板）
     *
     * @param {string} templateId - 模板 ID
     * @param {object} updates - 更新字段
     * @returns {object} 更新后的模板
     */
    updatePrivateTemplate(templateId, updates) {
      // 检查是否为公共模板
      const isPublic = PUBLIC_TEMPLATES.some(t => t.id === templateId)
      if (isPublic) {
        throw new Error('Cannot update public template')
      }

      const index = this.privateTemplates.findIndex(t => t.id === templateId)
      if (index === -1) {
        throw new Error(`Private template not found: ${templateId}`)
      }

      const template = this.privateTemplates[index]
      Object.assign(template, updates, {
        id: template.id,
        scope: 'private',
        userId: template.userId
      })

      this._persist()
      return template
    },

    /**
     * 删除自定义模板（拒绝公共模板）
     *
     * @param {string} templateId - 模板 ID
     */
    deletePrivateTemplate(templateId) {
      // 检查是否为公共模板
      const isPublic = PUBLIC_TEMPLATES.some(t => t.id === templateId)
      if (isPublic) {
        throw new Error('Cannot delete public template')
      }

      const index = this.privateTemplates.findIndex(t => t.id === templateId)
      if (index === -1) {
        throw new Error(`Private template not found: ${templateId}`)
      }

      this.privateTemplates.splice(index, 1)
      this._persist()
    }
  }
})
