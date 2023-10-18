package main

import (

	// sqlite
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// 实现版本号获取服务
// 通过http请求获取版本号

func main() {
	config, err := src.loadConfig("config.toml")
	if err != nil {
		panic(err)
	}

	// 0. 加载数据库
	db, err := sql.Open("sqlite3", config.DatabasePath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 1. 创建路由
	r := gin.Default()

	// 2. 绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/version", func(c *gin.Context) {
		// c.JSON：返回json格式的数据
		c.JSON(http.StatusOK, gin.H{
			"version": "1.0.0",
		})
	})

	// 3. 监听端口，默认在8080
	r.Run()
}

func LoadConfig(s string) {
	panic("unimplemented")
}
