package services_test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"web/service/gin/model/entities"
	mocks "web/service/gin/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var services mocks.MockIAlbumService

var expectedResult []entities.Album = []entities.Album{{
	ID:     uuid.MustParse("3fc7046e-666b-46a5-8028-b54f122118cf"),
	Title:  "Test_1",
	Artist: "TestA_1",
	Price:  10.0,
}, {
	ID:     uuid.MustParse("00000000-0000-0000-0000-000000000000"),
	Title:  "Test_2",
	Artist: "TestA_2",
	Price:  20.0,
}}

func TestGetAlbums(t *testing.T) {

	gin.SetMode(gin.TestMode)
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r := gin.Default()
	c, _ := gin.CreateTestContext(w)

	c.IndentedJSON(200, expectedResult)

	mockCtrl := gomock.NewController(t)
	mockClient := mocks.NewMockIAlbumService(mockCtrl)
	mockClient.EXPECT().GetAlbums().Return(expectedResult)

	r.GET("/albums", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, mockClient.GetAlbums())
	})

	// Check to see if the response was what you expected
	req, err := http.NewRequest(http.MethodGet, "/albums", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestGetAlbumById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.IndentedJSON(200, expectedResult[1])

	mockCtrl := gomock.NewController(t)
	mockClient := mocks.NewMockIAlbumService(mockCtrl)
	mockClient.EXPECT().GetAlbumById("00000000-0000-0000-0000-000000000000").Return(200, expectedResult[1])

	r.GET("/albums/:id", func(context *gin.Context) {
		status, body := mockClient.GetAlbumById("00000000-0000-0000-0000-000000000000")
		context.IndentedJSON(status, body)
	})

	// Define the method of request, URL and the body of request
	req, err := http.NewRequest(http.MethodGet, "/albums/00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestPostAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.IndentedJSON(201, expectedResult[1])

	mockCtrl := gomock.NewController(t)
	mockClient := mocks.NewMockIAlbumService(mockCtrl)
	mockClient.EXPECT().PostAlbums(expectedResult[1]).Return(201, expectedResult[1])

	r.POST("/albums", func(ctx *gin.Context) {
		var newAlbum entities.Album
		if err := ctx.BindJSON(&newAlbum); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, newAlbum)
		}
		status, body := mockClient.PostAlbums(newAlbum)
		ctx.IndentedJSON(status, body)
	})

	jsonValue, _ := json.Marshal(expectedResult[1])

	// Check to see if the response was what you expected
	req, err := http.NewRequest(http.MethodPost, "/albums", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusCreated, w.Code, "Correctly status code '201' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestPutAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.IndentedJSON(200, expectedResult[1])

	mockCtrl := gomock.NewController(t)
	mockClient := mocks.NewMockIAlbumService(mockCtrl)
	mockClient.EXPECT().PutAlbum(expectedResult[1]).Return(200, expectedResult[1])

	r.PUT("/albums", func(context *gin.Context) {
		var updatedAlbum entities.Album
		if err := context.BindJSON(&updatedAlbum); err != nil {
			context.IndentedJSON(http.StatusBadRequest, expectedResult[1])
		}
		status, body := mockClient.PutAlbum(updatedAlbum)
		context.IndentedJSON(status, body)
	})

	jsonValue, _ := json.Marshal(expectedResult[1])

	// Check to see if the response was what you expected
	req, err := http.NewRequest(http.MethodPut, "/albums", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusOK, w.Code, "Correctly status code '200' - OK")
	assert.NotNil(t, w.Body.String(), "Correctly body response - OK")
}

func TestDeleteAlbum(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.IndentedJSON(204, expectedResult[1])

	mockCtrl := gomock.NewController(t)
	mockClient := mocks.NewMockIAlbumService(mockCtrl)
	mockClient.EXPECT().DeleteAlbum("00000000-0000-0000-0000-000000000000").Return()

	r.DELETE("/albums/:id", func(context *gin.Context) {
		mockClient.DeleteAlbum("00000000-0000-0000-0000-000000000000")
		context.IndentedJSON(200, http.NoBody)
	})

	// Define the method of request, URL and the body of request
	req, err := http.NewRequest(http.MethodDelete, "/albums/00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	assert.Equal(t, http.StatusNoContent, w.Code, "Correctly status code '200' - OK")
}
