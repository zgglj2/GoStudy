package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "jrj")
		c.String(http.StatusOK, "hello "+name)
	})

	r.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert")

		username := c.PostForm("username")
		password := c.PostForm("password")

		hobby := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type: %s, username: %s, password: %s, hobby: %v", type1, username, password, hobby))
	})

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("upload: %s", file.Filename))
	})

	r.POST("/upload_multi", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get error: "+err.Error())
			return
		}
		files := form.File["files"]
		for _, file := range files {
			log.Println(file.Filename)
			err = c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				c.String(http.StatusBadRequest, "upload file error: "+err.Error())
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})

	v1 := r.Group("v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 login")
		})
	}
	v2 := r.Group("v2")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 login")
		})
	}
	r.Run(":8080")
}
