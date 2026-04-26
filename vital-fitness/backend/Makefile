.PHONY: build run clean swagger test

# 编译
build:
	go build -o main ./cmd/main.go

# 运行
run:
	go run ./cmd/main.go

# 清理
clean:
	rm -f main

# 生成 Swagger 文档
swagger:
	swag init -g cmd/main.go -o docs

# 运行测试
test:
	go test ./... -v

# 安装开发工具
tools:
	go install github.com/swaggo/swag/cmd/swag@latest
