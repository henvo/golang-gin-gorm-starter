# Golang Gin GORM starter

This repository contains a starter template for rapid microservice development
with Go. It uses [Gin](https://github.com/gin-gonic/gin) and
[GORM](https://gorm.io).

## Installation
* Get the repository from GitHub
``` bash
git clone git@github.com:henvo/golang-gin-gorm-starter.git
```
* Install dependencies
``` bash
go get
```

* Run app
```
go run main.go
```

## Things to consider
* Rename all instances of `github.com/henvo/golang-gin-gorm/` to your package
* Switch from the default database driver (mysql) to sqlite, postgresql, ...

## Env variables

* PORT (Default: `8080`)
* GIN_MODE (Default: `debug`)
* DATABASE_HOST (Default: `localhost:3306`)
* DATABASE_NAME (Default: `go_test`)
* DATABASE_USERNAME (Default: `golang`)
* DATABASE_PASSWORD (Default: `golang`)
