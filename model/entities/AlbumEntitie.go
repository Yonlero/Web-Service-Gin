package entities

import (
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
)

// -----------------------------Structs-------------------------
type Album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}

type VerifyI interface {
	checkTitle() bool
	checkArtist() bool
	checkPrice() bool
}

func (a Album) checkTitle() bool {
	return len(a.Title) != 0 && reflect.TypeOf(a.Title) != nil
}

func (a Album) checkArtist() bool {
	return len(a.Artist) != 0 && reflect.TypeOf(a.Artist) != nil
}

func (a Album) checkPrice() bool {
	return reflect.TypeOf(a.Price) != nil && a.Price > 0
}

func (a Album) CheckFields() bool {
	if a.checkArtist() && a.checkPrice() && a.checkTitle() {
		log.Println(a.printLogs())
		return true
	}
	return false
}

func (a Album) printLogs() string {
	return string("Album ID: " + a.ID.String() +
		"\nAlbum Title: " + a.Title + " - " + fmt.Sprintf("%v", a.checkTitle()) +
		"\nAlbum Artist: " + a.Artist + " - " + fmt.Sprintf("%v", a.checkArtist()) +
		"\nAlbum Price: " + fmt.Sprintf("%v", a.Price) + " - " + fmt.Sprintf("%v", a.checkPrice()))
}
