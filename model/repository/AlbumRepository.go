package repository

import (
	"net/http"
	"time"
	er "web/service/gin/errors"
	"web/service/gin/model/database"
	e "web/service/gin/model/entities"
)

type (
	Repository struct {
	}
	AlbumRepositoryI interface {
		GetAllAlbums() []e.Album
		GetAlbumById(id string) (*e.Album, *er.ErrorBodyResponse)
		CreateNewAlbum(newAlbum e.Album) (*int64, *er.ErrorBodyResponse)
		UpdateAlbum(updatedAlbum e.Album) (*int64, *er.ErrorBodyResponse)
		DeleteAlbum(id string)
	}
)

func (Repository) GetAllAlbums() []e.Album {
	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	rows, _ := db.Query("SELECT * FROM tb_albums")

	var albums []e.Album

	for rows.Next() {
		var alb e.Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		} else {
			albums = append(albums, alb)
		}
	}

	database.CloseConnection(db)
	return albums
}

func (Repository) GetAlbumById(id string) (*e.Album, *er.ErrorBodyResponse) {
	var alb e.Album

	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	rows := db.QueryRow("SELECT * FROM tb_albums WHERE id=$1", id)

	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		database.CloseConnection(db)
		return nil, createErrorBody(http.StatusBadRequest, "Album Not Found", http.ErrBodyNotAllowed)
	} else {
		database.CloseConnection(db)
		return &alb, nil
	}
}

func (r Repository) CreateNewAlbum(newAlbum e.Album) (*int64, *er.ErrorBodyResponse) {
	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	if !newAlbum.CheckFields() {
		return nil, createErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed)
	}

	rows, err := db.Exec("INSERT INTO tb_albums (title, artist, price) VALUES ($1, $2, $3)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, _ := rows.RowsAffected()
	return &rowsAffected, nil
}

func (r Repository) UpdateAlbum(updatedAlbum e.Album) (*int64, *er.ErrorBodyResponse) {
	db := database.OpenConnection()
	rows, err := db.Exec("UPDATE tb_albums SET title = $1, artist = $2, price = $3 WHERE id= $4", updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, updatedAlbum.ID)
	if err != nil {
		return nil, createErrorBody(http.StatusBadRequest, "Please fill the body correctly", http.ErrBodyNotAllowed)
	}
	rowsAffected, _ := rows.RowsAffected()
	return &rowsAffected, nil
}

func (r Repository) DeleteAlbum(id string) {
	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	_, err := db.Exec("DELETE FROM tb_albums WHERE id=$1;", id)
	if err != nil {
		return
	}
}

func createErrorBody(status int, msg string, err error) *er.ErrorBodyResponse {
	errorResponse := er.ErrorBodyResponse{Timestamp: time.Now(),
		Status:  status,
		Message: msg,
		Errors:  err}
	return &errorResponse
}
