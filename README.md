# Navigator 导航系统

## 项目简介

Navigator是一个基于Go语言开发的导航系统后端服务，提供了完整的导航功能支持，包括用户认证、权限管理、页面管理等核心功能。

## 技术栈

- **语言**: Go 1.25.1
- **Web框架**: Gin 1.11.0
- **ORM框架**: GORM 1.31.1
- **数据库**: MySQL
- **缓存**: Redis 8.11.5
- **日志**: Zap 1.27.1
- **认证**: JWT
- **配置管理**: TOML

## 项目结构

```
navigator/
├── api/                # API相关代码
│   ├── error/          # 错误定义
│   └── jwt/            # JWT认证
├── app/                # 应用入口
│   ├── config/         # 配置文件
│   └── main.go         # 主入口文件
├── conf/               # 配置管理
├── dao/                # 数据访问层
├── manager/            # 业务管理层
├── model/              # 数据模型
├── server/             # 服务器相关
│   └── http/           # HTTP服务器
├── service/            # 业务逻辑层
├── go.mod              # Go模块定义
├── go.sum              # 依赖版本锁定
├── build.sh            # 构建脚本
└── navigator.service   # 系统服务配置
```

## 快速开始

### 环境要求

- Go 1.25.1 或更高版本
- MySQL 5.7 或更高版本
- Redis 5.0 或更高版本

### 安装依赖

```bash
go mod download
```

### 配置文件

1. 复制配置文件模板

```bash
cp app/config/production/config.toml.example app/config/production/config.toml
```

2. 编辑配置文件，设置数据库连接、Redis连接等信息

### 运行项目

#### 开发环境

```bash
go run app/main.go
```

#### 生产环境

1. 构建项目

```bash
./build.sh
```

2. 启动服务

```bash
systemctl start navigator
```

## 核心功能

- **用户认证**: 基于JWT的用户认证系统
- **权限管理**: 基于RBAC的权限管理系统
- **页面管理**: 导航页面的创建、编辑、删除等操作
- **错误处理**: 统一的错误处理中间件
- **日志管理**: 基于Zap的结构化日志系统

## 部署

### 系统服务

项目提供了`navigator.service`文件，可以用于部署为系统服务。

1. 复制服务配置文件

```bash
cp navigator.service /etc/systemd/system/
```

2. 重载系统服务配置

```bash
systemctl daemon-reload
```

3. 启动服务

```bash
systemctl start navigator
```

4. 设置开机自启

```bash
systemctl enable navigator
```

### 容器部署

可以使用Docker容器化部署，需要编写Dockerfile和docker-compose.yml文件。

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开Pull Request

## 许可证

本项目采用MIT许可证 - 详情请参阅LICENSE文件
