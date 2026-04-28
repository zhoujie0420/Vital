# Vital Fitness 版本规划

## 当前进度

### ✅ 已完成

**基础设施**
- ✅ 数据库设计（7 张表：users, exercises, workouts, weight_records, mood_records, diet_records, foods）
- ✅ Docker Compose 编排（MySQL + API + Web）
- ✅ 后端配置管理（Viper，支持 YAML + 环境变量覆盖）
- ✅ 日志系统（zap）、数据库连接、GORM AutoMigrate
- ✅ CORS 中间件

**用户认证**
- ✅ JWT 中间件（生成 / 解析 / 鉴权）
- ✅ 微信小程序登录（code2session → 自动注册 → 签发 token）
- ✅ 用户资料查看 / 更新接口（白名单字段过滤）
- ✅ 前端 token 持久化 + 自动携带 + 401 自动跳转登录

**训练模块**
- ✅ Exercise 动作库查询 / 创建（handler → service → dao 全链路）
- ✅ Workout CRUD 完整实现（创建、列表、详情、更新、删除）
- ✅ 前端：Hevy/Strong 风格组表格记录、分页加载、骨架屏、详情编辑弹窗

**体重模块**
- ✅ WeightRecord 创建 / 列表 / 删除 + BMI 自动计算
- ✅ 前端：大号数字输入、BMI 实时计算、涨跌趋势标识、分页加载

**饮食记录模块**
- ✅ DietRecord CRUD + Food 食物库查询 / 自定义创建
- ✅ 前端：份量选择弹窗（快捷克数 + 实时营养计算）、今日热量概览 + 营养素进度条、按餐次分区

**心情模块**
- ✅ MoodRecord 创建 / 列表（service + dao 已实现）
- ✅ 前端：表情选择、标签多选、彩色分数圆点

**统计模块**
- ✅ 首页仪表盘 + 统计页接口（支持 week/month/year）
- ✅ 前端：深色活动环卡片、柱状图、体重趋势条形图

**饮食规划功能（spec: diet-plan）**
- ✅ 5 个公共模板定义（均衡饮食、碳循环、碳水渐降、高蛋白增肌、低碳减脂）
- ✅ 营养计算逻辑（`dietPlanCalc.js`）：固定比例计算、碳循环当日目标、碳水渐降周递减、剩余热量计算、输入验证
- ✅ Pinia Store（`dietPlan.js`）：方案 CRUD、模板 CRUD、激活/停用、localStorage 持久化/恢复
- ✅ 属性测试（fast-check）：模板聚合、宏量计算、输入验证、碳循环查找、碳水渐降计算、剩余热量等 8 个属性测试
- ✅ 单元测试：5 个模板具体计算验证、边界条件测试
- ✅ 页面注册（pages.json 已添加 plan 和 plan-edit）
- ✅ vitest + fast-check 测试框架配置

**前端设计系统**
- ✅ 统一 SCSS 设计系统（翡翠绿主色 + 暖灰中性色）
- ✅ 可复用 Mixins：card、primary-button、segment-control、skeleton-shimmer、stagger-fade 等
- ✅ 全局字体：SF Pro + PingFang SC
- ✅ Pinia 状态管理、请求工具封装

---

### 🔧 进行中（饮食规划 spec 剩余任务）

**Store 测试（Task 4.5-4.13）**
- ⏳ Store 属性测试：私有模板元数据、公共模板不可变、快照隔离、单一激活不变量、持久化往返、更新保留身份、删除清除激活
- ⏳ Store 单元测试：CRUD 操作状态变更验证

**方案列表页 UI（Task 6）**
- ⏳ `diet/plan.vue`：当前为空壳，需实现激活方案展示、用户方案列表、模板列表、创建/编辑/删除操作

**方案编辑页 UI（Task 7）**
- ⏳ `diet/plan-edit.vue`：当前为空壳，需实现模式路由（模板/自定义/编辑）、固定比例表单、碳循环周历配置、碳水渐降配置

**首页饮食模块（Task 8）**
- ⏳ `index/index.vue` 嵌入饮食规划进度：剩余热量、营养素进度条、快捷记录按钮、无方案引导

---

### ❌ 待完成（非 spec 范围）

**后端修复**
- ❌ 心情模块路由被注释（router.go），需取消注释启用

**数据初始化**
- ❌ 动作库（exercises）预置数据
- ❌ 食物库（foods）预置数据
- ❌ MySQL 初始化 SQL 脚本

**功能缺口**
- ❌ 心情记录删除接口
- ❌ 体重记录更新接口
- ❌ 饮水量记录
- ❌ 训练照片上传
- ❌ 目标体重设置
- ❌ 数据导出

**部署**
- ❌ Nginx SSL 配置
- ❌ Swagger 文档完善

**测试**
- ❌ 后端单元测试 / API 集成测试
