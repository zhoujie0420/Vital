# Vital Fitness - 身体状态记录系统

## 项目简介

Vital Fitness 是一个全栈健康管理微信小程序，帮助用户记录和追踪健身训练、饮食营养、体重变化和心理状态。采用前后端分离架构，支持微信小程序和 H5 双端访问。

## 技术栈

### 前端
- **框架**: Vue 3 + uni-app 3.x（微信小程序 + H5）
- **状态管理**: Pinia
- **构建工具**: Vite 5
- **样式**: SCSS 设计系统（翡翠绿主色调，健康感视觉风格）
- **包管理**: npm

### 后端
- **语言**: Go 1.25
- **Web 框架**: Gin
- **数据库**: MySQL 8.0
- **ORM**: GORM
- **认证**: JWT（golang-jwt/v5）
- **配置**: Viper（YAML + 环境变量）
- **日志**: zap
- **API 文档**: Swagger

### 部署
- **容器化**: Docker + Docker Compose
- **反向代理**: Nginx
- **域名**: api.bodysynclab.fun

## 项目结构

```
vital-fitness/
├── frontend/                   # 前端项目 (uni-app)
│   ├── src/
│   │   ├── pages/              # 页面（17个）
│   │   │   ├── auth/           # 登录、注册
│   │   │   ├── index/          # 首页仪表盘
│   │   │   ├── workout/        # 训练（列表、添加、详情）
│   │   │   ├── diet/           # 饮食（列表、添加、详情）
│   │   │   ├── mood/           # 心情（列表、添加）
│   │   │   ├── weight/         # 体重（列表、添加）
│   │   │   ├── statistics/     # 数据统计
│   │   │   └── user/           # 个人中心、资料、设置
│   │   ├── api/                # API 请求层（6 个模块）
│   │   ├── store/              # Pinia 状态管理
│   │   ├── styles/             # 设计系统（SCSS 变量 + Mixins）
│   │   ├── utils/              # 工具函数（请求封装、日期格式化）
│   │   └── static/             # 静态资源（tabbar 图标）
│   ├── package.json
│   ├── Dockerfile
│   └── nginx.conf
├── backend/                    # 后端项目 (Go + Gin)
│   ├── cmd/main.go             # 入口（AutoMigrate 7 张表）
│   ├── internal/
│   │   ├── config/             # 配置管理（Viper）
│   │   ├── router/             # 路由定义
│   │   ├── handler/            # 请求处理器（6 个）
│   │   ├── service/            # 业务逻辑层（6 个）
│   │   ├── dao/                # 数据访问层（5 个）
│   │   ├── model/              # 数据模型（6 个 + 请求/响应结构体）
│   │   ├── middleware/         # JWT 认证、CORS
│   │   └── utils/              # 数据库连接、日志
│   ├── configs/config.yaml
│   ├── go.mod / go.sum
│   └── Dockerfile
├── docker-compose.yml          # MySQL + API + Web
└── deploy.sh
```

## 功能模块

### 用户认证
- 微信小程序一键登录（code2session → 自动注册）
- JWT Token 认证（Bearer 方式）
- 个人资料管理（昵称、性别、身高、体重）

### 训练记录
- 动作库管理（按分类：胸/背/腿/肩/手臂/核心）
- Hevy/Strong 风格组表格记录（逐组填写重量和次数）
- 按日期分组展示、按分类筛选
- 训练详情编辑（重量/组数/次数并排输入）和删除
- 分页加载 + 骨架屏

### 饮食记录
- 四餐记录（早餐/午餐/晚餐/加餐）
- 食物库搜索（关键词 + 分类）+ 自定义食物
- 份量选择（快捷克数 + 自定义输入 + 实时营养计算）
- 今日热量概览 + 三大营养素进度条
- 按餐次分区展示（薄荷健康风格）

### 体重管理
- 大号数字输入 + ±0.1 步进 + 快捷选择
- BMI 自动计算（带状态标签：偏瘦/正常/超重/肥胖）
- 体重趋势列表（涨跌标识）

### 心情日记
- 心情评分（1-10 分 + 彩色分数圆点）
- 心情标签多选（开心、焦虑、放松等）
- 心情描述记录

### 数据统计
- 首页仪表盘（深色活动环卡片：训练/热量/体重）
- 训练次数柱状图（周/月/年切换）
- 体重趋势条形图
- 热量摄入柱状图 + 日均热量

### 饮食规划（开发中）
- 5 个系统预设模板：均衡饮食、碳循环、碳水渐降、高蛋白增肌、低碳减脂
- 三种方案类型：固定比例（fixed）、碳循环（carb_cycle）、碳水渐降（carb_taper）
- 基于体重自动计算每日营养目标（热量、蛋白质、碳水、脂肪）
- 碳循环：按星期区分高碳日（训练日）和低碳日（休息日）
- 碳水渐降：按周递减碳水摄入，最低 50g 保底
- 自定义模板创建和管理
- 首页饮食进度模块（剩余热量、营养素进度条、快捷记录）
- 数据通过 localStorage 持久化，后续迁移后端 API

## 快速开始

### 环境要求
- Node.js 18+
- Go 1.25+
- MySQL 8.0+
- Docker 20.0+（可选）

### 前端开发

```bash
cd vital-fitness/frontend
npm install

# 微信小程序开发（用微信开发者工具导入 dist/dev/mp-weixin）
npm run dev:mp-weixin

# H5 开发
npm run dev:h5
```

### 后端开发

```bash
cd vital-fitness/backend
# 修改 configs/config.yaml 中的数据库配置
go run cmd/main.go
```

### Docker 部署

```bash
cd vital-fitness
docker-compose up -d
```

## API 接口

### 认证（无需 Token）
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/v1/auth/wx-login | 微信登录 |
| POST | /api/v1/auth/logout | 登出 |

### 用户
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/users/profile | 获取个人资料 |
| PUT | /api/v1/users/profile | 更新个人资料 |

### 训练
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/exercises/ | 获取动作库 |
| POST | /api/v1/exercises/ | 创建动作 |
| POST | /api/v1/workouts/ | 创建训练记录 |
| GET | /api/v1/workouts/ | 获取训练列表（支持 category/page/page_size） |
| GET | /api/v1/workouts/:id | 获取训练详情 |
| PUT | /api/v1/workouts/:id | 更新训练记录 |
| DELETE | /api/v1/workouts/:id | 删除训练记录 |

### 饮食
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/foods/ | 搜索食物库（keyword/category） |
| POST | /api/v1/foods/ | 创建自定义食物 |
| POST | /api/v1/diets/ | 创建饮食记录 |
| GET | /api/v1/diets/ | 获取饮食列表 |
| PUT | /api/v1/diets/:id | 更新饮食记录 |
| DELETE | /api/v1/diets/:id | 删除饮食记录 |

### 体重
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/v1/weights/ | 创建体重记录 |
| GET | /api/v1/weights/ | 获取体重列表 |
| DELETE | /api/v1/weights/:id | 删除体重记录 |

### 心情（路由待启用）
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/v1/moods/ | 创建心情记录 |
| GET | /api/v1/moods/ | 获取心情列表 |

### 统计
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/dashboard | 首页仪表盘数据 |
| GET | /api/v1/stats?period=week | 统计图表数据（week/month/year） |

## 设计系统

项目使用统一的 SCSS 设计系统（`src/styles/variables.scss`）：

- **主色**: 翡翠绿 `#10B981`（健康、活力）
- **强调色**: 暖橙 `#F97316`（能量、动力）
- **背景**: 暖白 `#FAFAF9`
- **字体**: SF Pro + PingFang SC
- **动画**: fadeInUp 交错入场、骨架屏 shimmer、弹性缓动 `cubic-bezier(0.16, 1, 0.3, 1)`
- **组件 Mixins**: card、primary-button、segment-control、input-field、empty-state、skeleton-shimmer 等

## 许可证

MIT License
