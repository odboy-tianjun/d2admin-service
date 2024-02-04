# d2admin-service

#### 项目说明

```text
前端：https://gitee.com/d2-projects/d2-admin-start-kit

后端：从零开始。无依赖Casbin，无二次封装，适合学习。
```

### 文档地址

```text
# gin
https://gin-gonic.com/docs/quickstart
# gorm
https://gorm.io/zh_CN/docs/index.html
```

### 结构说明
```text
1、所有API接口必须在"src/infra/bind_api.go"上定义与"gin.HandlerFunc"的绑定关系
2、所有API接口必须在"system_router"表中定义
```

### golang版本

```text
go1.21.5 darwin/amd64
```

### 安装依赖并启动

```bash
go mod tidy
go run main.go
```

### 构建

```bash
go build -o d2admin-service
```