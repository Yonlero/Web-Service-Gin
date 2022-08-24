package repository

import (
	er "web/service/gin/errors"
	data "web/service/gin/model/database"
	e "web/service/gin/model/entities"
)

type (
	AlbumRepository  struct{ data.IDatabase }
	IAlbumRepository interface {
		GetAllAlbums() []e.Album
		GetAlbumById(id string) any
		CreateNewAlbum(newAlbum e.Album) any
		UpdateAlbum(updatedAlbum e.Album) any
		DeleteAlbum(id string)
	}
)

func (r AlbumRepository) GetAllAlbums() []e.Album {
	var albums []e.Album
	// Initialize DB Config (just testing connection)
	db := r.IDatabase.OpenConnection()
	rows, _ := db.Query("SELECT * FROM tb_albums")

	for rows.Next() {
		var alb e.Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		} else {
			albums = append(albums, alb)
		}
	}

	err := db.Close()
	if err != nil {
		return nil
	}
	return albums
}

func (r AlbumRepository) GetAlbumById(id string) any {
	var alb e.Album

	// Initialize DB Config (just testing connection)
	db := r.IDatabase.OpenConnection()
	rows := db.QueryRow("SELECT * FROM tb_albums WHERE id=$1", id)

	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		r.IDatabase.CloseConnection(db)
		return er.ErrorBodyResponse{}
	} else {
		r.IDatabase.CloseConnection(db)
		return alb
	}
}

func (r AlbumRepository) CreateNewAlbum(newAlbum e.Album) any {
	// Initialize DB Config (just testing connection)
	db := r.IDatabase.OpenConnection()

	if !newAlbum.CheckFields() {
		return &er.ErrorBodyResponse{}
	}

	_, err := db.Exec("INSERT INTO tb_albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		r.IDatabase.CloseConnection(db)
		return er.ErrorBodyResponse{}
	}

	r.IDatabase.CloseConnection(db)
	return newAlbum
}

func (r AlbumRepository) UpdateAlbum(updatedAlbum e.Album) any {
	db := r.IDatabase.OpenConnection()
	_, err := db.Exec("UPDATE tb_albums SET title = $1, artist = $2, price = $3 WHERE id= $4", updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, updatedAlbum.ID)
	if err != nil {
		return er.ErrorBodyResponse{}
	}
	r.IDatabase.CloseConnection(db)
	return updatedAlbum
}

func (r AlbumRepository) DeleteAlbum(id string) {
	// Initialize DB Config (just testing connection)
	db := r.IDatabase.OpenConnection()
	_, err := db.Exec("DELETE FROM tb_albums WHERE id=$1;", id)
	r.IDatabase.CloseConnection(db)
	if err != nil {
		return
	}
}
