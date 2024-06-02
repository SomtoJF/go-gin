package main

import (
	"github.com/somtojf/go-gin/initializers"
	"github.com/somtojf/go-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
