# gin-start

gin web 开发脚手架

## 目录结构

> 借鉴 [project-layout](https://github.com/golang-standards/project-layout)

## 开发

```bash
go run main.go # 直接启动服务
````

### 自动重启

使用 [air](https://github.com/cosmtrek/air/blob/master/README.md) 工具可以监听文件改动，自动重启服务，便于开发。

1. 安装 air 工具

```bash
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

2. 启动 air

项目自带 air 配置文件（.air.toml），只需要在项目根目录运行 air

```bash
air       # 全局安装启动
./bin/air # 指定目录安装启动
```

## 特性

### 配置

godotenv 库会加载 .env 配置文件到环境变量中，通过 `os.Getenv("SERVER_ADDR")` 获取。
.env 配置文件不应该包含在版本库中

### 日志

使用 [logrus](https://github.com/sirupsen/logrus) 日志工具

### 数据库

使用 [gorm](https://gorm.io) 作为 ORM 库，支持多种类型的数据库，本项目默认使用 Mysql。

```bash
go get -u gorm.io/gorm
```

#### Mysql

```bash
go get -u github.com/go-sql-driver/mysql
```