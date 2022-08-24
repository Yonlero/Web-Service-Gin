package main

import (
	"net/http"
	"web/service/gin/model/database"
	"web/service/gin/model/entities"
	"web/service/gin/model/repository"
	"web/service/gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//Get default settings
	router := gin.Default()
	albumDatabase := database.Database{}
	albumRepository := repository.AlbumRepository{IDatabase: albumDatabase}
	service := services.AlbumService{IAlbumRepository: albumRepository}

	//Define Routes and functions
	router.GET("/albums", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, service.GetAlbums())
	})
	router.GET("/albums/:id", func(context *gin.Context) {
		status, body := service.GetAlbumById(context.Param("id"))
		context.IndentedJSON(status, body)
	})
	router.POST("/albums", func(context *gin.Context) {
		var newAlbum entities.Album
		if err := context.BindJSON(&newAlbum); err != nil {
			context.IndentedJSON(http.StatusBadRequest, services.CreateErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed))
		}
		status, body := service.PostAlbums(newAlbum)
		context.IndentedJSON(status, body)
	})
	router.PUT("/albums", func(context *gin.Context) {
		var updatedAlbum entities.Album
		if err := context.BindJSON(&updatedAlbum); err != nil {
			context.IndentedJSON(http.StatusBadRequest, services.CreateErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed))
		}
		status, body := service.PutAlbum(updatedAlbum)
		context.IndentedJSON(status, body)
	})
	router.DELETE("albums/:id", func(context *gin.Context) {
		service.DeleteAlbum(context.Param("id"))
		context.IndentedJSON(200, http.NoBody)
	})

	//Run the API
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
