/**
 * 版本变更日志
 * Changelog data
 *
 * 新增版本时请在数组首位添加记录，保持时间倒序。
 * type 取值：feature(新增) / improvement(优化) / fix(修复) / security(安全)
 */

export const CURRENT_VERSION = '1.1.0'

export const CHANGELOG = [
  {
    version: '1.1.0',
    date: '2026-05-06',
    title: '游客浏览体验',
    changes: [
      { type: 'improvement', text: '未登录用户可自由浏览首页、训练、饮食、体重、统计等功能' },
      { type: 'improvement', text: '仅在保存记录等写操作时才引导登录，不再强制登录' },
      { type: 'feature', text: '设置中新增版本变动查看入口' }
    ]
  },
  {
    version: '1.0.0',
    date: '2026-05-05',
    title: 'AI 食物识别',
    changes: [
      { type: 'feature', text: '饮食记录支持拍照 AI 识别，自动填充食物与份量' },
      { type: 'improvement', text: '识别结果支持按项勾选、调整份量后批量添加' },
      { type: 'improvement', text: '部署流程强制 --no-cache 确保每次部署最新代码' }
    ]
  },
  {
    version: '0.9.0',
    date: '2026-04-28',
    title: 'UI 改版与饮食规划',
    changes: [
      { type: 'feature', text: '全新 UI 设计系统：翡翠绿主色 + 暖灰中性色' },
      { type: 'feature', text: '饮食规划：5 套公共模板（均衡/碳循环/碳水渐降/高蛋白/低碳）' },
      { type: 'feature', text: '支持自定义方案，localStorage 持久化与单一激活约束' },
      { type: 'improvement', text: '训练记录改为 Hevy/Strong 风格组表格' }
    ]
  },
  {
    version: '0.8.0',
    date: '2026-04-26',
    title: '安全与工程化',
    changes: [
      { type: 'security', text: '移除 phpMyAdmin，避免数据库暴露' },
      { type: 'improvement', text: '后端补全 CRUD 接口，增加分页与参数校验' },
      { type: 'fix', text: 'Go 模块代理切换至 goproxy.cn，修复国内构建超时' }
    ]
  },
  {
    version: '0.7.0',
    date: '2026-04-06',
    title: '技术栈升级',
    changes: [
      { type: 'improvement', text: '升级到 uni-app 3.x + Vue 3 Composition API' },
      { type: 'improvement', text: '状态管理迁移到 Pinia' },
      { type: 'feature', text: '首页仪表盘 + 统计页 + 用户资料全部对接真实 API' }
    ]
  },
  {
    version: '0.6.0',
    date: '2026-04-05',
    title: '饮食记录模块',
    changes: [
      { type: 'feature', text: '食物库 + 饮食 CRUD 接口与前端对接' },
      { type: 'feature', text: 'TabBar 新增饮食入口' },
      { type: 'improvement', text: '份量选择弹窗支持快捷克数与实时营养计算' }
    ]
  },
  {
    version: '0.5.0',
    date: '2026-04-05',
    title: '体重与心情模块',
    changes: [
      { type: 'feature', text: '体重记录创建/列表/删除，自动计算 BMI' },
      { type: 'feature', text: '心情记录创建与列表展示' }
    ]
  },
  {
    version: '0.4.0',
    date: '2026-04-05',
    title: '训练记录模块',
    changes: [
      { type: 'feature', text: '动作库与训练记录完整 CRUD' },
      { type: 'improvement', text: '训练列表支持分页加载与骨架屏' },
      { type: 'fix', text: '训练记录去掉 Feeling 必填校验' }
    ]
  }
]

/**
 * 类型显示配置
 */
export const TYPE_LABELS = {
  feature: { text: '新增', color: 'feature' },
  improvement: { text: '优化', color: 'improvement' },
  fix: { text: '修复', color: 'fix' },
  security: { text: '安全', color: 'security' }
}

export default {
  CURRENT_VERSION,
  CHANGELOG,
  TYPE_LABELS
}
