package services

import (
	"log"
	"net/http"
	"web/service/gin/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// -----------------------------Variables-------------------------
// I'm not using uuid.New() because I'd like use the same ID in the tests
var albums = []entities.Album{
	{ID: uuid.MustParse("b69ff5c3-253b-4fd2-9568-293300552dee"), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: uuid.MustParse("fb04cb76-8095-48c7-9536-1bb72a332689"), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: uuid.MustParse("0a01e3d9-f977-4458-8d57-1c0af687a58f"), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// -----------------------------Functions-------------------------
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Print(id.String())
		panic(err)
	}
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
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

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbum(c *gin.Context) {
	var id uuid.UUID = uuid.MustParse(c.Param("id"))

	for index, a := range albums {
		if a.ID == id {
			albums = removeIndex(albums, index)
		}
	}

	if c.Errors != nil {
		c.Errors.JSON()
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
	c.IndentedJSON(http.StatusOK, albums)

}

func removeIndex(s []entities.Album, index int) []entities.Album {
	return append(s[:index], s[index+1:]...)
}
