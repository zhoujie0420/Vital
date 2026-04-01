#!/bin/bash

# Vital Fitness 启动脚本

echo "🚀 正在启动 Vital Fitness 应用..."

# 检查Docker是否安装
if ! command -v docker &> /dev/null
then
    echo "❌ 未检测到 Docker，请先安装 Docker"
    exit 1
fi

# 检查Docker Compose是否安装
if ! command -v docker-compose &> /dev/null
then
    echo "❌ 未检测到 Docker Compose，请先安装 Docker Compose"
    exit 1
fi

# 构建并启动所有服务
echo "🔨 正在构建并启动所有服务..."
docker-compose up -d

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 10

# 检查服务状态
echo "🔍 检查服务状态..."
docker-compose ps

echo ""
echo "✅ 应用启动完成！"
echo ""
echo "应用查看: http://localhost"
echo "API文档: http://localhost:8080/swagger/index.html"
echo "数据库管理: http://localhost:8081"
echo "Redis管理: http://localhost:8082"
echo ""
echo "🔧 停止服务请运行: docker-compose down"