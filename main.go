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
	r.GET("/posts/:id", controllers.GetOnePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.GET("/posts", controllers.GetManyPosts)
	r.Run() // listen and serve on 0.0.0.0:8080
}
