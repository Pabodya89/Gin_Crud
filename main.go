package main

import (
	"Gin_Crud/controller"
	"Gin_Crud/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
  router := gin.Default()
  
  router.POST("/posts", controller.PostsCreate)
  router.GET("/posts", controller.PostsIndex)
  router.GET("/posts/:id", controller.PostShow)
  router.PUT("/posts/:id", controller.PostUpdate)
  router.DELETE("/posts/:id", controller.PostDelete)

  router.Run()
}