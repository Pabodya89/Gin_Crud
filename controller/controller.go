package controller

import (
	"Gin_Crud/initializers"
	"Gin_Crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body string
		Title string 
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.db.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post" : post,
	})

}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts models.Post
	initializers.db.Find(&posts)
}

func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.db.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id
	id := c.Param("id")
	//Get the data of req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//find the post were updating
	var post models.Post
	initializers.db.First(&post, id)

	//Update it
	initializers.db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Response it
	c.JSON(200, gin.H{
		"post" : post,
	})
}

func PostDelete(c *gin.Context)  {
	//Get the id of the url
	id := c.Param("id")

	//delete the posts
	initializers.db.Delete(&models.Post{}, id)

	//response it
	c.Status(200)
}