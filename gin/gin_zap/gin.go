package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config := LogConfig{
		Level:      "info",
		Filename:   "./gin.log",
		MaxSize:    10,
		MaxBackups: 2,
	}

	err := InitLogger(&config)
	if err != nil {
		fmt.Println("InitLogger failed, err: ", err)
		return
	}

	r := gin.Default()
	r.Use(GinLogger(Logger), GinRecovery(Logger, true))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
