#!/bin/bash
# Vital Fitness 一键部署脚本
# 使用方法: 在服务器上执行
# curl -sSL https://raw.githubusercontent.com/zhoujie0420/Vital/main/vital-fitness/deploy.sh | bash

set -e

echo "========================================="
echo "  Vital Fitness 部署脚本"
echo "========================================="

# 1. 安装 Docker（如果未安装）
if ! command -v docker &> /dev/null; then
    echo ">>> 安装 Docker..."
    curl -fsSL https://get.docker.com | sh
    systemctl start docker
    systemctl enable docker
    echo ">>> Docker 安装完成"
else
    echo ">>> Docker 已安装，跳过"
fi

# 2. 安装 Docker Compose（如果未安装）
if ! docker compose version &> /dev/null; then
    echo ">>> 安装 Docker Compose 插件..."
    apt-get update && apt-get install -y docker-compose-plugin 2>/dev/null || \
    yum install -y docker-compose-plugin 2>/dev/null || true
    echo ">>> Docker Compose 安装完成"
else
    echo ">>> Docker Compose 已安装，跳过"
fi

# 3. 克隆项目
PROJECT_DIR="/opt/vital-fitness"
if [ -d "$PROJECT_DIR" ]; then
    echo ">>> 更新项目代码..."
    cd "$PROJECT_DIR"
    git pull origin main
else
    echo ">>> 克隆项目..."
    git clone https://github.com/zhoujie0420/Vital.git "$PROJECT_DIR"
fi

# 4. 启动服务
cd "$PROJECT_DIR/vital-fitness"
echo ">>> 停止旧服务（如果有）..."
docker compose down 2>/dev/null || true

echo ">>> 构建并启动服务..."
docker compose up -d --build

# 5. 等待服务启动
echo ">>> 等待服务启动..."
sleep 10

# 6. 检查服务状态
echo ""
echo "========================================="
echo "  部署完成！服务状态："
echo "========================================="
docker compose ps

echo ""
echo "访问地址："
echo "  前端:     http://$(curl -s ifconfig.me):80"
echo "  后端API:  http://$(curl -s ifconfig.me):8080/health"
echo "  phpMyAdmin: http://$(curl -s ifconfig.me):8081"
echo "========================================="
