package main

import (
	"web/service/gin/model/repository"
	"web/service/gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//Get default settings
	router := gin.Default()
	albumRepository := repository.Repository{}
	service := services.AlbumService{AlbumRepositoryI: albumRepository}

	//Define Routes and functions
	router.GET("/albums", service.GetAlbums)
	router.GET("/albums/:id", service.GetAlbumById)
	router.POST("/albums", service.PostAlbums)
	router.PUT("/albums", service.PutAlbum)
	router.DELETE("albums/:id", service.DeleteAlbum)

	//Run the API
	router.Run("localhost:8080")
}
