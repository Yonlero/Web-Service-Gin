package services

import (
	"fmt"
	"log"
	"net/http"
	"web/service/gin/model/database"
	"web/service/gin/model/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// -----------------------------Functions-------------------------
func GetAlbums(c *gin.Context) {
	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	rows, err := db.Query("SELECT * FROM tb_albums")
	if err != nil {
		panic(err.Error())
	}

	var albums []entities.Album

	for rows.Next() {
		var alb entities.Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		} else {
			albums = append(albums, alb)
		}
	}

	c.IndentedJSON(http.StatusOK, albums)
	database.CloseConnection(db)
}

func GetAlbumById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Fatalln("ID PROBLEM: " + id.String())
		panic(err)
	}

	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	rows, err := db.Query("SELECT * FROM tb_albums")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var alb entities.Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		}

		if alb.ID == id {
			c.IndentedJSON(http.StatusOK, alb)
			break
		}
	}

	if c.Errors != nil {
		c.Errors.JSON()
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func PostAlbums(c *gin.Context) {
	var newAlbum entities.Album
	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.CheckFields()

	rows, err := db.Exec("INSERT INTO tb_albums (title, artist, price) VALUES ('" + newAlbum.Title + "','" + newAlbum.Artist + "'," + fmt.Sprintf("%f", newAlbum.Price) + ");")
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, _ := rows.RowsAffected()

	log.Println("Rows affected: " + fmt.Sprintf("%d", rowsAffected))

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func PutAlbum(c *gin.Context) {
	var updatedAlbum entities.Album
	db := database.OpenConnection()
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	updatedAlbum.CheckFields()

	rows, err := db.Exec("UPDATE tb_albums SET title = '" + updatedAlbum.Title + "', artist = '" + updatedAlbum.Artist + "', price = " + fmt.Sprintf("%f", updatedAlbum.Price) + "  WHERE id='" + updatedAlbum.ID.String() + "';")
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, _ := rows.RowsAffected()

	log.Println("Rows affected: " + fmt.Sprintf("%d", rowsAffected))
	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

func DeleteAlbum(c *gin.Context) {
	var id uuid.UUID = uuid.MustParse(c.Param("id"))

	// Initialize DB Config (just testing connection)
	db := database.OpenConnection()
	rows, err := db.Exec("DELETE FROM tb_albums WHERE id='" + id.String() + "';")
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, _ := rows.RowsAffected()

	log.Println("Rows affected: " + fmt.Sprintf("%d", rowsAffected))
	c.IndentedJSON(http.StatusNoContent, nil)

	if c.Errors != nil {
		c.Errors.JSON()
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

}
