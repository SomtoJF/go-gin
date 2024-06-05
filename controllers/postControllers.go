package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/somtojf/go-gin/initializers"
	"github.com/somtojf/go-gin/models"
)

type DataBody struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type DataBodyOptional struct {
	Title string
	Body  string
}

type GetOnePostBody struct {
	ID int `json:"id" binding:"required"`
}

func PostCreate(c *gin.Context) {
	var body DataBody

	// Use ShouldBindJSON to get more detailed error information
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func GetOnePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func GetManyPosts(c *gin.Context) {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var body DataBodyOptional

	// Get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Use ShouldBindJSON to get more detailed error information
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(http.StatusCreated, gin.H{"data": post})

}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusNoContent, gin.H{})
}
