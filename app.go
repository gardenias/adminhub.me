package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/**")

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})

	engine.GET("/hello", func(context *gin.Context) {
		context.HTML(200, "hello.html",gin.H{"username":"Riemann"})
	})

	engine.Run()
}
