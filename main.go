package main

import (
	"github.com/gin-gonic/gin"
	"github.com/somtojf/go-gin/controllers"
	"github.com/somtojf/go-gin/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
