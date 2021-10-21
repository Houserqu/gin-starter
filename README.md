# gin-start

gin web 开发脚手架，根据个人开发经验封装的一些基础能力，并提供一些接入 mysql、redis 等中间件的示例。

基础架构代码在 master 分支，示例代码在独立分支，如下

- [mysql](https://github.com/Houserqu/gin-starter/tree/mysql): 使用 gorm 接入 mysql，并提供 增删改查示例。

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

使用 [viper](https://github.com/spf13/viper) 库来管理系统配置，同时开启 watch 模式，默认加载项目根目录的 config.yaml 文件
config.yaml 配置文件不应该包含在版本库中

### 日志

使用 [logrus](https://github.com/sirupsen/logrus) 日志工具，支持记录请求ID，用于查看日志链路。

```golang
import 	"houserqu.com/gin-starter/internal"

internal.Log(c).Info("123") // 由于无法获取协程上下文，所以需要显式的传递 gin.Context，才能记录 request id
```

### 数据库

使用 [gorm](https://gorm.io) 作为 ORM 库，支持多种类型的数据库，本项目默认使用 Mysql。

```bash
go get -u gorm.io/gorm
```

#### Mysql

```bash
go get -u github.com/go-sql-driver/mysql
```

## 部署

### Docker 部署

1. 构建镜像
```bash
docker build -t gin-starter:latest .
```

2. 启动容器
```bash
docker run -d --env-file prod.env -p 8090:8088 -v /Users/houserqu/gin-starter/logs:/app/logs gin-starter:latest
```

-p: 指定映射端口，容器内服务端口可以通过 env 文件配置  
-v: 指定日志文件主机挂载目录，容器内日志目录可以通过 env 文件配置

### 直接部署

```bash
go build main.go # 构建
./main           # 直接启动
pm2 start ./main # pm 守护进程启动
```