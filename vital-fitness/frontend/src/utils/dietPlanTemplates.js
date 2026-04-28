/**
 * 饮食方案模板定义
 * Diet plan template definitions
 *
 * 5 个公共模板，每个模板包含名称、描述、类型和基于体重的计算参数。
 * 公共模板 scope='public'，用户不可编辑或删除。
 */

/**
 * @typedef {Object} CalcConfigFixed
 * @property {number} caloriesPerKg - 每公斤体重热量
 * @property {number} [proteinPerKg] - 每公斤体重蛋白质（克）
 * @property {number} [proteinRatio] - 蛋白质占总热量比例（0-1）
 * @property {number} [carbsRatio] - 碳水占总热量比例（0-1）
 * @property {number|null} [fatRatio] - 脂肪占总热量比例（0-1），null 表示由剩余热量推算
 */

/**
 * @typedef {Object} CalcConfigCarbCycle
 * @property {{ caloriesPerKg: number, proteinPerKg: number, carbsPerKg: number, fatPerKg: number }} high - 高碳日参数
 * @property {{ caloriesPerKg: number, proteinPerKg: number, carbsPerKg: number, fatPerKg: number }} low - 低碳日参数
 */

/**
 * @typedef {Object} CalcConfigCarbTaper
 * @property {number} caloriesPerKg - 每公斤体重热量
 * @property {number} proteinPerKg - 每公斤体重蛋白质（克）
 * @property {number} initialCarbsPerKg - 每公斤体重初始碳水（克）
 * @property {number} fatPerKg - 每公斤体重脂肪（克）
 */

/**
 * @typedef {Object} DietPlanTemplate
 * @property {string} id - 唯一标识
 * @property {'public'|'private'} scope - 模板范围
 * @property {string} name - 模板名称
 * @property {string} description - 模板描述
 * @property {'fixed'|'carb_cycle'|'carb_taper'} type - 方案类型
 * @property {CalcConfigFixed|CalcConfigCarbCycle|CalcConfigCarbTaper} calcConfig - 计算配置
 * @property {Record<number, 'high'|'low'>} [defaultCycleConfig] - 碳循环默认周配置
 * @property {{ totalWeeks: number, weeklyReduction: number }} [defaultTaperConfig] - 碳水渐降默认配置
 */

/** @type {DietPlanTemplate[]} */
export const PUBLIC_TEMPLATES = [
  {
    id: 'tpl_balanced',
    scope: 'public',
    name: '均衡饮食',
    description: '40%碳水 / 30%蛋白 / 30%脂肪，适合日常健康饮食',
    type: 'fixed',
    calcConfig: {
      caloriesPerKg: 30,
      proteinRatio: 0.30,
      carbsRatio: 0.40,
      fatRatio: 0.30
    }
  },
  {
    id: 'tpl_carb_cycle',
    scope: 'public',
    name: '碳循环',
    description: '训练日高碳、休息日低碳，适合增肌减脂交替',
    type: 'carb_cycle',
    calcConfig: {
      high: { caloriesPerKg: 33, proteinPerKg: 2.0, carbsPerKg: 4.0, fatPerKg: 0.8 },
      low: { caloriesPerKg: 26, proteinPerKg: 2.2, carbsPerKg: 1.5, fatPerKg: 1.0 }
    },
    defaultCycleConfig: { 0: 'low', 1: 'high', 2: 'low', 3: 'high', 4: 'low', 5: 'high', 6: 'low' }
  },
  {
    id: 'tpl_carb_taper',
    scope: 'public',
    name: '碳水渐降',
    description: '每周递减碳水摄入，适合备赛或短期减脂',
    type: 'carb_taper',
    calcConfig: {
      caloriesPerKg: 28,
      proteinPerKg: 2.0,
      initialCarbsPerKg: 3.0,
      fatPerKg: 0.9
    },
    defaultTaperConfig: { totalWeeks: 8, weeklyReduction: 15 }
  },
  {
    id: 'tpl_high_protein',
    scope: 'public',
    name: '高蛋白增肌',
    description: '蛋白质2.0g/kg体重，适合增肌期',
    type: 'fixed',
    calcConfig: {
      caloriesPerKg: 35,
      proteinPerKg: 2.0,
      carbsRatio: 0.45,
      fatRatio: null
    }
  },
  {
    id: 'tpl_low_carb',
    scope: 'public',
    name: '低碳减脂',
    description: '低碳水高蛋白，适合减脂期',
    type: 'fixed',
    calcConfig: {
      caloriesPerKg: 24,
      proteinPerKg: 2.0,
      carbsRatio: 0.20,
      fatRatio: null
    }
  }
]

export default {
  PUBLIC_TEMPLATES
}
