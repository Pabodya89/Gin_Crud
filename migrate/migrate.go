package main

import (
	"Gin_Crud/initializers"
	"Gin_Crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main()  {
	initializers.db.AutoMigrate(&models.Post{})
}