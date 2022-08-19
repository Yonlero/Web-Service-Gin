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

/*
	---------------------------Functions------------------------
*/
// GetAlbums Return all albums in DB
func (s AlbumService) GetAlbums(c *gin.Context) {
	var albums []entities.Album = s.AlbumRepositoryI.GetAllAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func (s AlbumService) GetAlbumById(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	response, err := s.AlbumRepositoryI.GetAlbumById(id.String())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (s AlbumService) PostAlbums(c *gin.Context) {
	var newAlbum entities.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, createErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed))
	}
	response, err := s.AlbumRepositoryI.CreateNewAlbum(newAlbum)
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
		c.IndentedJSON(http.StatusBadRequest, createErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed))
		return
	}
	response, err := s.AlbumRepositoryI.UpdateAlbum(updatedAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (s AlbumService) DeleteAlbum(c *gin.Context) {
	var id uuid.UUID = uuid.MustParse(c.Param("id"))
	s.AlbumRepositoryI.DeleteAlbum(id.String())
	c.IndentedJSON(http.StatusNoContent, nil)
}

func createErrorBody(status int, msg string, err error) e.ErrorBodyResponse {
	errorResponse := e.ErrorBodyResponse{Timestamp: time.Now(),
		Status:  status,
		Message: msg,
		Errors:  err}
	return errorResponse
}
