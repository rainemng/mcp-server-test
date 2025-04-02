# MCP Server Test

这是一个用于测试 MCP Server 的项目。

## 项目结构

```
.
├── README.md
├── .gitignore
├── cmd/            # 主要的应用程序入口
├── internal/       # 私有应用程序和库代码
├── pkg/           # 可以被外部应用程序使用的库代码
└── api/           # API 协议定义和 swagger/OpenAPI 规范
```

## 开发环境要求

- Go 1.20 或更高版本
- Docker
- Make

## 快速开始

1. 克隆仓库
```bash
git clone https://github.com/rainemng/mcp-server-test.git
cd mcp-server-test
```

2. 安装依赖
```bash
go mod download
```

3. 运行服务
```bash
go run cmd/main.go
```

## 测试

运行所有测试：
```bash
go test ./...
```

## 构建

构建二进制文件：
```bash
go build -o bin/server cmd/main.go
```

## 贡献

欢迎提交 Pull Request 和 Issue！

## 许可证

MIT License