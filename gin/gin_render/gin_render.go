package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "abc"})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name": "glj"})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}

		c.ProtoBuf(http.StatusOK, data)
	})

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.html")
	r.GET("/someHTML", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "glj"})
	})

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/async", func(c *gin.Context) {
		ctx_copy := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("异步执行: ", ctx_copy.Request.URL.Path)
		}()
	})

	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("同步执行: ", c.Request.URL.Path)
	})
	r.Run(":8080")
}
