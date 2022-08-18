package services

import (
	"net/http"
	"time"
	e "web/service/gin/errors"
	"web/service/gin/model/entities"
	r "web/service/gin/model/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AlbumService  struct{ r.AlbumRepositoryI }
	AlbumServiceI interface {
		GetAlbums(c *gin.Context)
		GetAlbumById(c *gin.Context)
		PostAlbums(c *gin.Context)
		PutAlbum(c *gin.Context)
		DeleteAlbum(c *gin.Context)
	}
)

var repository r.Repository

/*
	---------------------------Functions------------------------
*/
// GetAlbums Return all albums in DB
func (s AlbumService) GetAlbums(c *gin.Context) {
	var albums []entities.Album = repository.GetAllAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func (s AlbumService) GetAlbumById(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	response, err := repository.GetAlbumById(id.String())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (s AlbumService) PostAlbums(c *gin.Context) {
	var newAlbum entities.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		errorResponse := e.ErrorBodyResponse{Timestamp: time.Now(),
			Status:  http.StatusBadRequest,
			Message: "Cannot read your body request",
			Errors:  http.ErrBodyNotAllowed}
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
	}
	response, err := repository.CreateNewAlbum(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, response)
}

func (s AlbumService) PutAlbum(c *gin.Context) {
	var updatedAlbum entities.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	if !updatedAlbum.CheckFields() {
		errorResponse := e.ErrorBodyResponse{Timestamp: time.Now(),
			Status:  http.StatusBadRequest,
			Message: "Please fill the body correctly",
			Errors:  http.ErrBodyNotAllowed}
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	response, err := repository.UpdateAlbum(updatedAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (s AlbumService) DeleteAlbum(c *gin.Context) {
	var id uuid.UUID = uuid.MustParse(c.Param("id"))
	repository.DeleteAlbum(id.String())
	c.IndentedJSON(http.StatusNoContent, nil)
}
