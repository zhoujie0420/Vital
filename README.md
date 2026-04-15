# Vital Fitness - 身体状态记录系统

## 项目概述

Vital Fitness 是一个多用户的身体状态记录系统，旨在帮助用户追踪和管理他们的健身进度、心理状态、体重变化和饮食情况。该系统采用前后端分离架构，支持多端访问（小程序、H5）。

## 技术架构

### 前端技术栈
- **框架**: uni-app (支持编译到微信小程序、H5等多个平台)
- **UI库**: uView UI (专为uni-app设计的UI框架)
- **状态管理**: Vuex
- **网络请求**: axios/flyio
- **本地存储**: localStorage/uni.storage

### 后端技术栈
- **语言**: Go
- **Web框架**: Gin
- **数据库**: MySQL 8.0+
- **ORM**: GORM
- **API文档**: Swagger
- **认证授权**: JWT
- **缓存**: Redis
- **日志**: zap

### 部署架构
- **容器化**: Docker + Docker Compose
- **反向代理**: Nginx
- **云服务**: 阿里云/腾讯云
- **对象存储**: 用于存储用户头像等静态资源

## 功能模块

### 1. 用户系统
- 用户注册/登录 (手机号/邮箱)
- 第三方登录 (微信授权登录)
- 用户个人信息管理
- 用户权限管理

### 2. 健身记录模块 (核心功能)
- 训练计划制定与管理
- 动作库维护 (深蹲、卧推、硬拉等)
- 每次训练记录：
  - 训练日期和时间
  - 训练动作选择
  - 重量、组数、次数记录
  - 训练感受评分
  - 训练照片上传
- 训练数据分析与可视化
- 个人最佳成绩(PB)记录

### 3. 心理状态模块
- 每日心情打卡 (1-10分评分)
- 心情标签选择 (兴奋、疲惫、焦虑等)
- 心理状态文字描述
- 心情趋势图表展示
- 压力指数跟踪

### 4. 体重管理模块
- 每日体重记录
- BMI自动计算
- 体脂率记录 (可选)
- 体重变化趋势图
- 目标体重设置与进度追踪

### 5. 饮食记录模块
- 三餐记录 (早餐、午餐、晚餐、加餐)
- 食物数据库 (常见食物热量)
- 自定义食物添加
- 卡路里自动计算
- 营养成分分析 (蛋白质、碳水、脂肪)
- 饮水记录

## API接口设计

### 认证接口
- POST /api/v1/auth/register - 用户注册
- POST /api/v1/auth/login - 用户登录
- POST /api/v1/auth/logout - 用户登出

### 用户接口
- GET /api/v1/users/profile - 获取用户信息
- PUT /api/v1/users/profile - 更新用户信息

### 健身接口
- GET /api/v1/exercises - 获取动作列表
- POST /api/v1/workouts - 创建训练记录
- GET /api/v1/workouts - 获取训练记录列表
- GET /api/v1/workouts/:id - 获取单条训练记录
- PUT /api/v1/workouts/:id - 更新训练记录
- DELETE /api/v1/workouts/:id - 删除训练记录

### 心情接口
- POST /api/v1/moods - 创建心情记录
- GET /api/v1/moods - 获取心情记录列表

### 体重接口
- POST /api/v1/weights - 创建体重记录
- GET /api/v1/weights - 获取体重记录列表

### 饮食接口
- POST /api/v1/diets - 创建饮食记录
- GET /api/v1/diets - 获取饮食记录列表

## 项目目录结构

```
vital-fitness/
├── frontend/                 # 前端项目
│   ├── pages/                # 页面目录
│   ├── components/           # 组件目录
│   ├── store/                # Vuex状态管理
│   ├── utils/                # 工具函数
│   └── api/                  # API接口封装
├── backend/                  # 后端项目
│   ├── cmd/                  # 应用入口
│   ├── internal/             # 内部代码
│   │   ├── handler/          # 请求处理器
│   │   ├── service/          # 业务逻辑层
│   │   ├── dao/              # 数据访问层
│   │   ├── model/            # 数据模型
│   │   └── middleware/       # 中间件
│   ├── pkg/                  # 公共包
│   ├── configs/              # 配置文件
│   └── docs/                 # 文档
├── docker/                   # Docker配置
└── README.md                 # 项目说明文档
```

## 部署方案

### 开发环境
- 使用docker-compose一键启动所有服务
- 热重载开发模式

### 生产环境
- Nginx反向代理
- HTTPS加密传输
- 数据库主从复制
- Redis缓存集群
- 日志收集与监控

## 项目特点

1. **多端支持**: 一套代码多端运行 (微信小程序、H5)
2. **数据安全**: JWT认证，敏感信息加密存储
3. **易于扩展**: 微服务友好架构设计
4. **性能优化**: Redis缓存，数据库索引优化
5. **用户体验**: 响应式设计，流畅交互体验