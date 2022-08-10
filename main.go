package main

import (
	"web/service/gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//Get default settings
	router := gin.Default()

	//Define Routes and functions
	router.GET("/albums", services.GetAlbums)
	router.GET("/albums/:id", services.GetAlbumById)
	router.POST("/albums", services.PostAlbums)
	router.DELETE("albums/:id", services.DeleteAlbum)

	//Run the API
	router.Run("localhost:8080")
}
