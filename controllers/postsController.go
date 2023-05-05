package controllers

import (
	"github.com/gin-gonic/gin"
	"main.go/initializers"
	"main.go/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Get data off req body

	//Create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)
}

func PostsShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with tem
	c.JSON(200, gin.H{
		"posts": post,
	})
}
func PostsUpdate(c *gin.Context) {

	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//Responde with it
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Delete(&models.Post{}, id)
	//Responde with it
	c.JSON(200, gin.H{
		"message": post,
	})
}
