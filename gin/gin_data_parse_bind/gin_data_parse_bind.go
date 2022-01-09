package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" bingding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" bingding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		err := c.ShouldBindJSON(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		err := c.Bind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.GET("/login/:user/:password", func(c *gin.Context) {
		var login Login
		c.ShouldBindUri(&login)

		c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s", login.User, login.Password))
	})
	r.Run(":8080")

}
