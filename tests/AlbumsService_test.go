package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"web/service/gin/model/entities"
	"web/service/gin/services"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/albums", services.GetAlbums)

	// Check to see if the response was what you expected
	req, err := http.NewRequest(http.MethodGet, "/albums", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestGetAlbumById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/albums/:id", services.GetAlbumById)

	// Define the method of request, URL and the body of request
	req, err := http.NewRequest(http.MethodGet, "/albums/b69ff5c3-253b-4fd2-9568-293300552dee", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestPostAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/albums", services.PostAlbums)

	newAlbum := entities.Album{
		ID:     uuid.New(),
		Title:  "Test Title",
		Artist: "Test Artist",
		Price:  56.99,
	}
	jsonValue, _ := json.Marshal(newAlbum)

	// Check to see if the response was what you expected
	req, err := http.NewRequest(http.MethodPost, "/albums", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusCreated, w.Code, "Correctly status code '201' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestPutAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/albums", services.PutAlbum)

	updatedAlbum := entities.Album{
		ID:     uuid.MustParse("b69ff5c3-253b-4fd2-9568-293300552dee"),
		Title:  "Test Title Updated",
		Artist: "Test Artist Updated",
		Price:  56.99,
	}

	jsonValue, _ := json.Marshal(updatedAlbum)

	// Check to see if the response was what you expected
	req, err := http.NewRequest(http.MethodPut, "/albums", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestDeleteAlbum(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/albums/:id", services.DeleteAlbum)

	// Define the method of request, URL and the body of request
	req, err := http.NewRequest(http.MethodDelete, "/albums/b69ff5c3-253b-4fd2-9568-293300552dee", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
}
