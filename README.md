# Vital Fitness - 身体状态记录系统

## 项目简介

Vital Fitness 是一个全栈健康管理微信小程序，帮助用户记录和追踪健身训练、饮食营养、体重变化。采用前后端分离架构，支持微信小程序和 H5 双端访问。

## 技术栈

### 前端
- **框架**: Vue 3 + uni-app 3.x（微信小程序 + H5）
- **状态管理**: Pinia
- **构建工具**: Vite 5
- **样式**: SCSS 设计系统（翡翠绿主色调）
- **测试**: Vitest + fast-check（属性测试）

### 后端
- **语言**: Go 1.25
- **Web 框架**: Gin
- **数据库**: MySQL 8.0
- **ORM**: GORM（AutoMigrate 7 张表）
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
├── frontend/                   # 前端项目 (uni-app + Vue 3)
│   ├── src/
│   │   ├── pages/              # 页面
│   │   │   ├── auth/           # 登录、注册
│   │   │   ├── index/          # 首页仪表盘
│   │   │   ├── workout/        # 训练（列表、添加、详情）
│   │   │   ├── diet/           # 饮食（列表、添加、详情、规划）
│   │   │   ├── weight/         # 体重（列表、添加）
│   │   │   ├── statistics/     # 数据统计
│   │   │   └── user/           # 个人中心、资料、设置
│   │   ├── api/                # API 请求层
│   │   ├── store/              # Pinia 状态管理
│   │   ├── styles/             # SCSS 设计系统
│   │   ├── utils/              # 工具函数
│   │   └── static/             # 静态资源
│   ├── package.json
│   └── Dockerfile
├── backend/                    # 后端项目 (Go + Gin)
│   ├── cmd/main.go             # 入口
│   ├── internal/
│   │   ├── config/             # 配置管理
│   │   ├── router/             # 路由定义
│   │   ├── handler/            # 请求处理器（6 个）
│   │   ├── service/            # 业务逻辑层（6 个）
│   │   ├── dao/                # 数据访问层（5 个）
│   │   ├── model/              # 数据模型（6 个）
│   │   ├── middleware/         # JWT 认证、CORS
│   │   └── utils/              # 数据库连接、日志
│   ├── configs/config.yaml.example
│   ├── Makefile
│   └── Dockerfile
├── mysql/init/                 # 数据库初始化 SQL
├── docker-compose.yml
├── deploy.sh                   # 一键部署脚本
└── start.sh                    # 本地启动脚本
```

## 功能模块

### 用户认证
- 微信小程序一键登录（code2session → 自动注册）
- JWT Token 认证
- 个人资料管理（昵称、性别、身高、体重）
- 退出登录

### 训练记录
- 动作库管理（按分类：胸/背/腿/肩/手臂/核心，预置 14 个动作）
- 组表格记录（逐组填写重量和次数）
- 按日期分组展示、按分类筛选
- 完整 CRUD + 分页加载 + 下拉刷新

### 饮食记录
- 四餐记录（早餐/午餐/晚餐/加餐）
- 食物库搜索 + 自定义食物
- 份量选择 + 实时营养计算
- 今日热量概览 + 三大营养素进度条
- 完整 CRUD + 分页加载 + 下拉刷新

### 饮食规划（开发中）
- 5 个系统预设模板（均衡饮食、碳循环、碳水渐降、高蛋白增肌、低碳减脂）
- 基于体重自动计算每日营养目标
- 碳循环 / 碳水渐降等多种方案类型
- Pinia Store + localStorage 持久化

### 体重管理
- 大号数字输入 + BMI 自动计算
- 体重趋势列表（涨跌标识）
- 创建 / 列表 / 删除 + 分页加载 + 下拉刷新

### 心情日记（路由暂停，代码已实现）
- 心情评分（1-10 分）+ 标签多选
- 心情描述记录

### 数据统计
- 首页仪表盘（深色活动环卡片：训练/热量/体重）
- 训练次数柱状图
- 体重趋势条形图
- 热量摄入柱状图
- 支持周/月/年时间段切换

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

# 微信小程序开发
npm run dev:mp-weixin
# 用微信开发者工具导入 dist/dev/mp-weixin

# H5 开发
npm run dev:h5
```

### 后端开发

```bash
cd vital-fitness/backend

# 复制配置文件并修改
cp configs/config.yaml.example configs/config.yaml
# 编辑 config.yaml 填入数据库和微信配置

# 运行
make run

# 或直接
go run cmd/main.go
```

### Docker 一键部署

```bash
cd vital-fitness
docker-compose up -d
```

### 服务器部署

```bash
curl -sSL https://raw.githubusercontent.com/zhoujie0420/Vital/main/vital-fitness/deploy.sh | bash
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
| GET | /api/v1/workouts/ | 获取训练列表（category/page/page_size） |
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

### 心情（路由暂停）
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/v1/moods/ | 创建心情记录 |
| GET | /api/v1/moods/ | 获取心情列表 |

### 统计
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/dashboard | 首页仪表盘数据 |
| GET | /api/v1/stats?period=week | 统计数据（week/month/year） |

## 设计系统

统一的 SCSS 设计系统（`src/styles/variables.scss`）：

- **主色**: 翡翠绿 `#10B981`
- **强调色**: 暖橙 `#F97316`
- **背景**: 暖白 `#FAFAF9`
- **字体**: SF Pro + PingFang SC
- **组件 Mixins**: card、primary-button、segment-control、skeleton-shimmer 等

## 许可证

MIT License
