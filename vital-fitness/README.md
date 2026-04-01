# Vital Fitness - 身体状态记录系统

## 项目简介

Vital Fitness 是一个全栈的健康管理应用，帮助用户记录和追踪健身训练、心理状态、体重变化和饮食情况。采用前后端分离架构，支持多端访问。

## 技术栈

### 前端 (uni-app)
- 框架: uni-app (支持微信小程序、H5等多端)
- UI库: uView UI
- 状态管理: Vuex
- 构建工具: Vite
- 包管理: npm

### 后端 (Go)
- 语言: Go 1.25
- Web框架: Gin
- 数据库: MySQL 8.0
- ORM: GORM
- 缓存: Redis 7
- API文档: Swagger
- 日志: zap

### 部署
- 容器化: Docker + Docker Compose
- 反向代理: Nginx
- 数据库管理: phpMyAdmin
- 缓存管理: Redis Commander

## 项目结构

```
vital-fitness/
├── frontend/                 # 前端项目 (uni-app)
│   ├── pages/                # 页面目录
│   │   ├── auth/             # 认证相关页面
│   │   ├── index/            # 首页
│   │   ├── workout/          # 训练记录页面
│   │   ├── mood/             # 心情记录页面
│   │   ├── weight/           # 体重记录页面
│   │   ├── diet/             # 饮食记录页面
│   │   ├── statistics/       # 统计数据页面
│   │   └── user/             # 用户中心页面
│   ├── components/           # 组件目录
│   ├── store/                # Vuex状态管理
│   ├── utils/                # 工具函数
│   ├── static/               # 静态资源
│   ├── App.vue               # 应用入口
│   ├── main.js               # 主配置文件
│   ├── package.json          # 依赖配置
│   ├── Dockerfile            # 前端Docker配置
│   └── nginx.conf            # Nginx配置
├── backend/                  # 后端项目 (Go)
│   ├── cmd/                  # 应用入口
│   ├── internal/             # 内部代码
│   │   ├── config/           # 配置管理
│   │   ├── router/           # 路由配置
│   │   ├── handler/          # 请求处理器
│   │   ├── service/          # 业务逻辑层
│   │   ├── dao/              # 数据访问层
│   │   ├── model/            # 数据模型
│   │   └── utils/            # 工具函数
│   ├── configs/              # 配置文件
│   ├── docs/                 # 文档
│   ├── deploy/               # 部署脚本
│   ├── go.mod                # Go模块文件
│   ├── go.sum                # Go依赖校验
│   └── Dockerfile            # 后端Docker配置
├── mysql/                    # 数据库初始化脚本
│   └── init/                 # 初始化SQL脚本
├── docker-compose.yml        # Docker编排配置
└── README.md                 # 项目说明文档
```

## 功能模块

### 1. 用户系统
- 用户注册/登录
- 第三方登录 (微信、QQ、Apple)
- 个人资料管理
- 隐私设置

### 2. 健身记录
- 训练动作库管理
- 每次训练详细记录 (重量、组数、次数)
- 训练感受评分
- 训练照片上传

### 3. 心理状态
- 每日心情打卡 (1-10分)
- 心情标签标记
- 心情描述记录
- 心情趋势图表

### 4. 体重管理
- 每日体重记录
- BMI自动计算
- 体重变化趋势图
- 目标体重设置

### 5. 饮食记录
- 三餐记录 (早餐、午餐、晚餐、加餐)
- 食物数据库查询
- 卡路里自动计算
- 营养成分分析
- 饮水记录

### 6. 数据统计
- 综合数据概览
- 训练趋势分析
- 体重变化图表
- 营养摄入分析
- 自定义时间范围统计

## 快速开始

### 环境要求
- Docker 20.0+
- Docker Compose 1.29+

### 启动项目

1. 克隆项目代码
```bash
git clone <repository-url>
cd vital-fitness
```

2. 启动所有服务
```bash
docker-compose up -d
```

3. 访问应用
- 前端页面: http://localhost
- API文档: http://localhost:8080/swagger/index.html
- 数据库管理: http://localhost:8081
- Redis管理: http://localhost:8082

### 开发环境

#### 后端开发
```bash
cd backend
go run cmd/main.go
```

#### 前端开发
```bash
cd frontend
npm install
npm run dev:h5
```

## API接口

### 认证接口
- POST /api/v1/auth/register - 用户注册
- POST /api/v1/auth/login - 用户登录
- POST /api/v1/auth/logout - 用户登出

### 用户接口
- GET /api/v1/users/profile - 获取用户信息
- PUT /api/v1/users/profile - 更新用户信息

### 训练接口
- GET /api/v1/exercises - 获取动作列表
- POST /api/v1/workouts - 创建训练记录
- GET /api/v1/workouts - 获取训练记录列表

### 心情接口
- POST /api/v1/moods - 创建心情记录
- GET /api/v1/moods - 获取心情记录列表

### 体重接口
- POST /api/v1/weights - 创建体重记录
- GET /api/v1/weights - 获取体重记录列表

### 饮食接口
- POST /api/v1/diets - 创建饮食记录
- GET /api/v1/diets - 获取饮食记录列表

## 部署说明

### 生产环境部署
1. 修改 docker-compose.yml 中的环境变量
2. 配置域名和SSL证书
3. 启动服务
```bash
docker-compose -f docker-compose.yml up -d
```

### 数据备份
```bash
# 备份数据库
docker exec vital_mysql mysqldump -u root -p vital_fitness > backup.sql

# 恢复数据库
docker exec -i vital_mysql mysql -u root -p vital_fitness < backup.sql
```

## 项目特点

1. **多端支持**: 一套代码编译到多个平台
2. **现代化技术栈**: Go + Vue.js + Docker
3. **完整的功能体系**: 覆盖健康管理的各个方面
4. **良好的用户体验**: 响应式设计，流畅交互
5. **易于部署**: Docker容器化，一键部署
6. **数据安全**: JWT认证，数据加密存储

## 开发计划

### 已完成功能
- ✅ 项目架构设计
- ✅ 前后端基础框架搭建
- ✅ 数据库设计与初始化
- ✅ 核心功能页面开发
- ✅ Docker容器化配置

### 待开发功能
- 🔧 后端API接口实现
- 🔧 前端数据对接
- 🔧 用户认证授权
- 🔧 数据统计图表
- 🔧 测试用例编写
- 📱 小程序端适配
- 📊 数据导出功能

## 贡献指南

欢迎提交 Issue 和 Pull Request 来改进项目！

## 许可证

MIT License