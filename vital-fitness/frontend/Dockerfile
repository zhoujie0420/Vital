# 使用Node.js镜像作为构建环境
FROM node:16-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm ci

# 复制源代码
COPY . .

# 构建H5版本
RUN npm run build:h5

# 使用Nginx作为生产环境
FROM nginx:alpine

# 从builder阶段复制构建产物
COPY --from=builder /app/dist/build/h5 /usr/share/nginx/html

# 复制Nginx配置
COPY nginx.conf /etc/nginx/nginx.conf

# 暴露端口
EXPOSE 80

# 启动Nginx
CMD ["nginx", "-g", "daemon off;"]