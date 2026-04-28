/**
 * 饮食方案计算工具
 * Diet plan calculation utilities
 *
 * 纯函数模块，负责所有营养目标计算逻辑。
 * 包括：模板宏量计算、碳循环日目标、碳水渐降周目标、剩余计算、输入验证。
 */

/**
 * 将数值四舍五入到 1 位小数
 * @param {number} value
 * @returns {number}
 */
function round1(value) {
  return Math.round(value * 10) / 10
}

/**
 * 确保值非负
 * @param {number} value
 * @returns {number}
 */
function nonNeg(value) {
  return Math.max(0, value)
}

/**
 * 根据模板和体重计算宏量营养素目标
 *
 * 支持三种模板类型：
 * - ratio-based (均衡饮食): 使用 caloriesPerKg 和比例计算
 * - per-kg (高蛋白增肌, 低碳减脂): 使用 proteinPerKg，剩余热量按比例分配
 * - carb_cycle: 分别计算高碳日和低碳日目标，返回 { high, low }
 * - carb_taper: 计算初始目标
 *
 * @param {import('./dietPlanTemplates').DietPlanTemplate} template - 模板对象
 * @param {number} bodyWeight - 用户体重（公斤）
 * @returns {{ calories: number, protein: number, carbs: number, fat: number } | { high: { calories: number, protein: number, carbs: number, fat: number }, low: { calories: number, protein: number, carbs: number, fat: number } }}
 */
export function calculateMacrosFromTemplate(template, bodyWeight) {
  if (!template || !template.calcConfig || !bodyWeight || bodyWeight <= 0) {
    return { calories: 0, protein: 0, carbs: 0, fat: 0 }
  }

  const { type, calcConfig } = template

  if (type === 'carb_cycle') {
    return calculateCarbCycleMacros(calcConfig, bodyWeight)
  }

  if (type === 'carb_taper') {
    return calculateCarbTaperMacros(calcConfig, bodyWeight)
  }

  // fixed type
  return calculateFixedMacros(calcConfig, bodyWeight)
}

/**
 * 计算固定类型模板的宏量营养素
 *
 * 支持两种计算模式：
 * 1. Ratio-based (均衡饮食): 所有宏量从热量比例推算
 *    - calories = bodyWeight * caloriesPerKg
 *    - protein = calories * proteinRatio / 4
 *    - carbs = calories * carbsRatio / 4
 *    - fat = calories * fatRatio / 9
 *
 * 2. Per-kg based (高蛋白增肌, 低碳减脂): 蛋白质按体重计算，碳水按比例，脂肪填充剩余
 *    - calories = bodyWeight * caloriesPerKg
 *    - protein = bodyWeight * proteinPerKg
 *    - carbs = calories * carbsRatio / 4
 *    - fat = (calories - protein*4 - carbs*4) / 9  (when fatRatio is null)
 *
 * 所有值四舍五入到整数且非负。
 *
 * @param {object} calcConfig
 * @param {number} bodyWeight
 * @returns {{ calories: number, protein: number, carbs: number, fat: number }}
 */
function calculateFixedMacros(calcConfig, bodyWeight) {
  const calories = Math.round(nonNeg(bodyWeight * calcConfig.caloriesPerKg))

  // Ratio-based: 均衡饮食 — all macros derived from calorie ratios
  if (calcConfig.proteinRatio !== undefined && calcConfig.proteinRatio !== null) {
    const protein = Math.round(nonNeg((calories * calcConfig.proteinRatio) / 4))
    const carbs = Math.round(nonNeg((calories * calcConfig.carbsRatio) / 4))
    const fat = Math.round(nonNeg((calories * calcConfig.fatRatio) / 9))
    return { calories, protein, carbs, fat }
  }

  // Per-kg based: 高蛋白增肌, 低碳减脂 — protein from per-kg, then derive remaining
  const protein = Math.round(nonNeg(bodyWeight * calcConfig.proteinPerKg))
  const proteinCalories = protein * 4

  // Carbs from ratio of total calories
  const carbs = Math.round(nonNeg((calories * calcConfig.carbsRatio) / 4))
  const carbsCalories = carbs * 4

  // Fat fills the remaining calories (fatRatio is null)
  const remainingCalories = Math.max(0, calories - proteinCalories - carbsCalories)
  const fat = Math.round(nonNeg(remainingCalories / 9))

  return { calories, protein, carbs, fat }
}

/**
 * 计算碳循环模板的高碳日和低碳日宏量营养素
 * @param {object} calcConfig - 包含 high 和 low 配置
 * @param {number} bodyWeight
 * @returns {{ high: { calories: number, protein: number, carbs: number, fat: number }, low: { calories: number, protein: number, carbs: number, fat: number } }}
 */
function calculateCarbCycleMacros(calcConfig, bodyWeight) {
  const calcDay = (dayConfig) => ({
    calories: Math.round(nonNeg(bodyWeight * dayConfig.caloriesPerKg)),
    protein: Math.round(nonNeg(bodyWeight * dayConfig.proteinPerKg)),
    carbs: Math.round(nonNeg(bodyWeight * dayConfig.carbsPerKg)),
    fat: Math.round(nonNeg(bodyWeight * dayConfig.fatPerKg))
  })

  return {
    high: calcDay(calcConfig.high),
    low: calcDay(calcConfig.low)
  }
}

/**
 * 计算碳水渐降模板的初始宏量营养素目标
 * @param {object} calcConfig
 * @param {number} bodyWeight
 * @returns {{ calories: number, protein: number, carbs: number, fat: number }}
 */
function calculateCarbTaperMacros(calcConfig, bodyWeight) {
  const calories = Math.round(nonNeg(bodyWeight * calcConfig.caloriesPerKg))
  const protein = Math.round(nonNeg(bodyWeight * calcConfig.proteinPerKg))
  const carbs = Math.round(nonNeg(bodyWeight * calcConfig.initialCarbsPerKg))
  const fat = Math.round(nonNeg(bodyWeight * calcConfig.fatPerKg))
  return { calories, protein, carbs, fat }
}

/**
 * 统一获取方案当日目标（调度器）
 *
 * 根据 plan.type 分发到对应计算函数：
 * - fixed: 直接返回 plan.macroTargets
 * - carb_cycle: 委托给 getCarbCycleDayTargets
 * - carb_taper: 委托给 getCarbTaperDayTargets
 *
 * @param {import('../store/dietPlan').UserDietPlan} plan - 用户饮食方案
 * @param {Date} [date=new Date()] - 目标日期
 * @returns {{ calories: number, protein: number, carbs: number, fat: number }}
 */
export function getTodayTargets(plan, date = new Date()) {
  if (!plan) {
    return { calories: 0, protein: 0, carbs: 0, fat: 0 }
  }

  switch (plan.type) {
    case 'fixed':
      return { ...plan.macroTargets }
    case 'carb_cycle':
      return getCarbCycleDayTargets(plan, date)
    case 'carb_taper':
      return getCarbTaperDayTargets(plan, date)
    default:
      return { ...plan.macroTargets }
  }
}

/**
 * 获取碳循环方案当日目标
 *
 * 根据 date 的星期几（0=周日..6=周六）查找 cycleConfig.dayMap，
 * 返回对应的 highDayTargets 或 lowDayTargets。
 *
 * @param {object} plan - 碳循环方案，包含 cycleConfig
 * @param {Date} date - 目标日期
 * @returns {{ calories: number, protein: number, carbs: number, fat: number }}
 */
export function getCarbCycleDayTargets(plan, date) {
  if (!plan || !plan.cycleConfig) {
    return { calories: 0, protein: 0, carbs: 0, fat: 0 }
  }

  const dayOfWeek = date.getDay() // 0=Sunday..6=Saturday
  const dayType = plan.cycleConfig.dayMap[dayOfWeek]

  if (dayType === 'high') {
    return { ...plan.cycleConfig.highDayTargets }
  }
  return { ...plan.cycleConfig.lowDayTargets }
}

/**
 * 获取碳水渐降方案当日目标
 *
 * 计算逻辑：
 * 1. weekNumber = floor((date - startedAt) / (7 * 24 * 60 * 60 * 1000)) + 1
 * 2. currentCarbs = max(initialCarbs - (weekNumber - 1) * weeklyReduction, minCarbs || 50)
 * 3. 返回调整后的 macroTargets，碳水替换为 currentCarbs
 *
 * @param {object} plan - 碳水渐降方案，包含 taperConfig 和 macroTargets
 * @param {Date} date - 目标日期
 * @returns {{ calories: number, protein: number, carbs: number, fat: number }}
 */
export function getCarbTaperDayTargets(plan, date) {
  if (!plan || !plan.taperConfig || !plan.macroTargets) {
    return { calories: 0, protein: 0, carbs: 0, fat: 0 }
  }

  const { taperConfig, macroTargets } = plan
  const startedAt = new Date(plan.startedAt)
  const msPerWeek = 7 * 24 * 60 * 60 * 1000

  const weekNumber = Math.floor((date - startedAt) / msPerWeek) + 1
  const minCarbs = taperConfig.minCarbs || 50
  const currentCarbs = round1(
    Math.max(taperConfig.initialCarbs - (weekNumber - 1) * taperConfig.weeklyReduction, minCarbs)
  )

  return {
    calories: macroTargets.calories,
    protein: macroTargets.protein,
    carbs: currentCarbs,
    fat: macroTargets.fat
  }
}

/**
 * 计算剩余宏量营养素
 *
 * 对每个宏量营养素执行简单减法：targets.X - consumed.X
 *
 * @param {{ calories: number, protein: number, carbs: number, fat: number }} targets - 目标值
 * @param {{ calories: number, protein: number, carbs: number, fat: number }} consumed - 已摄入值
 * @returns {{ calories: number, protein: number, carbs: number, fat: number }}
 */
export function calculateRemaining(targets, consumed) {
  if (!targets || !consumed) {
    return { calories: 0, protein: 0, carbs: 0, fat: 0 }
  }

  return {
    calories: round1(targets.calories - consumed.calories),
    protein: round1(targets.protein - consumed.protein),
    carbs: round1(targets.carbs - consumed.carbs),
    fat: round1(targets.fat - consumed.fat)
  }
}

/**
 * 验证方案输入数据
 *
 * 规则：
 * - calories 必须为正数（> 0）
 * - protein, carbs, fat 必须为非负数（>= 0）
 *
 * @param {{ calories: number, protein: number, carbs: number, fat: number }} data - 输入数据
 * @returns {{ valid: boolean, errors: Record<string, string> }}
 */
export function validatePlanInput(data) {
  const errors = {}

  if (data.calories === undefined || data.calories === null || typeof data.calories !== 'number' || isNaN(data.calories) || data.calories <= 0) {
    errors.calories = '热量目标必须为大于0的数字'
  }

  if (data.protein === undefined || data.protein === null || typeof data.protein !== 'number' || isNaN(data.protein) || data.protein < 0) {
    errors.protein = '蛋白质目标必须为非负数'
  }

  if (data.carbs === undefined || data.carbs === null || typeof data.carbs !== 'number' || isNaN(data.carbs) || data.carbs < 0) {
    errors.carbs = '碳水目标必须为非负数'
  }

  if (data.fat === undefined || data.fat === null || typeof data.fat !== 'number' || isNaN(data.fat) || data.fat < 0) {
    errors.fat = '脂肪目标必须为非负数'
  }

  return {
    valid: Object.keys(errors).length === 0,
    errors
  }
}

/**
 * 验证碳循环配置
 *
 * 规则：dayMap 中必须至少包含一个 'high' 和一个 'low'
 *
 * @param {Record<number, 'high'|'low'>} dayMap - 周配置映射（0=周日..6=周六）
 * @returns {{ valid: boolean, errors: Record<string, string> }}
 */
export function validateCycleConfig(dayMap) {
  const errors = {}

  if (!dayMap || typeof dayMap !== 'object') {
    errors.dayMap = '碳循环配置不能为空'
    return { valid: false, errors }
  }

  const values = Object.values(dayMap)
  const hasHigh = values.includes('high')
  const hasLow = values.includes('low')

  if (!hasHigh) {
    errors.dayMap = '碳循环配置至少需要一个高碳日'
  } else if (!hasLow) {
    errors.dayMap = '碳循环配置至少需要一个低碳日'
  }

  if (!hasHigh && !hasLow) {
    errors.dayMap = '碳循环配置至少需要一个高碳日和一个低碳日'
  }

  return {
    valid: Object.keys(errors).length === 0,
    errors
  }
}

/**
 * 验证碳水渐降配置
 *
 * 规则：
 * - totalWeeks 必须为正整数
 * - weeklyReduction 必须为正数
 *
 * @param {{ totalWeeks: number, weeklyReduction: number }} config - 渐降配置
 * @returns {{ valid: boolean, errors: Record<string, string> }}
 */
export function validateTaperConfig(config) {
  const errors = {}

  if (!config || typeof config !== 'object') {
    errors.config = '碳水渐降配置不能为空'
    return { valid: false, errors }
  }

  if (
    config.totalWeeks === undefined ||
    config.totalWeeks === null ||
    typeof config.totalWeeks !== 'number' ||
    isNaN(config.totalWeeks) ||
    !Number.isInteger(config.totalWeeks) ||
    config.totalWeeks <= 0
  ) {
    errors.totalWeeks = '总周数必须为正整数'
  }

  if (
    config.weeklyReduction === undefined ||
    config.weeklyReduction === null ||
    typeof config.weeklyReduction !== 'number' ||
    isNaN(config.weeklyReduction) ||
    config.weeklyReduction <= 0
  ) {
    errors.weeklyReduction = '每周递减量必须为正数'
  }

  return {
    valid: Object.keys(errors).length === 0,
    errors
  }
}
