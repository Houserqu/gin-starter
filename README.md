# gin-start

gin web 开发脚手架

## 目录结构

> 参考 [project-layout](https://github.com/golang-standards/project-layout)

## 开发规范

## 特性

- [gin-start](#gin-start)
  - [目录结构](#目录结构)
  - [开发规范](#开发规范)
  - [特性](#特性)
    - [配置](#配置)
    - [日志](#日志)
    - [数据库](#数据库)
      - [Mysql](#mysql)

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