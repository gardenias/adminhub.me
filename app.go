package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


type Auth struct {
	Username string
	Token string
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
}

func main() {
	auth := Auth{Username:viper.GetString("username"),Token:viper.GetString("token")}

	engine := gin.Default()

	engine.LoadHTMLGlob("template/**")

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})

	engine.GET("/hello", func(context *gin.Context) {
		context.HTML(200, "hello.html",gin.H{"username":"Riemann"})
	})

	engine.GET("/user", func(context *gin.Context) {
		//context.JSON(200,gin.con)
	})

	engine.Run()
}
