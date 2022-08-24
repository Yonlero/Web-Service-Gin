package services

import (
	"log"
	"net/http"
	"time"
	er "web/service/gin/errors"
	"web/service/gin/model/entities"
	r "web/service/gin/model/repository"
)

type (
	AlbumService  struct{ r.IAlbumRepository }
	IAlbumService interface {
		GetAlbums() []entities.Album
		GetAlbumById(id string) (int, any)
		PostAlbums(album entities.Album) (int, any)
		PutAlbum(album entities.Album) (int, any)
		DeleteAlbum(id string)
	}
)

/*
	---------------------------Functions------------------------
*/
// GetAlbums Return all albums in DB
func (s AlbumService) GetAlbums() []entities.Album {
	return s.IAlbumRepository.GetAllAlbums()
}

func (s AlbumService) GetAlbumById(id string) (int, any) {
	response := s.IAlbumRepository.GetAlbumById(id)

	switch a := response.(type) {
	default:
		log.Printf("GetById request type: %T", a)
		return http.StatusBadRequest, CreateErrorBody(http.StatusNotFound, "Id not found", http.ErrAbortHandler)
	case entities.Album:
		return http.StatusOK, response
	}
}

func (s AlbumService) PostAlbums(album entities.Album) (int, any) {
	if !album.CheckFields() {
		return http.StatusBadRequest, CreateErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed)
	}

	response := s.IAlbumRepository.CreateNewAlbum(album)
	switch a := response.(type) {
	default:
		log.Printf("PostAlbums request type: %T", a)
		return http.StatusBadRequest, response
	case entities.Album:
		return http.StatusOK, response
	}
}

func (s AlbumService) PutAlbum(album entities.Album) (int, any) {
	if !album.CheckFields() {
		return http.StatusBadRequest, CreateErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed)
	}

	response := s.IAlbumRepository.UpdateAlbum(album)
	switch a := response.(type) {
	default:
		log.Printf("PutAlbum request type: %T", a)
		return http.StatusBadRequest, response
	case entities.Album:
		return http.StatusOK, response
	}
}

func (s AlbumService) DeleteAlbum(id string) {
	s.IAlbumRepository.DeleteAlbum(id)
}

func CreateErrorBody(status int, msg string, err error) er.ErrorBodyResponse {
	errorResponse := er.ErrorBodyResponse{
		Timestamp: time.Now(),
		Status:    status,
		Message:   msg,
		Errors:    err,
	}
	return errorResponse
}
