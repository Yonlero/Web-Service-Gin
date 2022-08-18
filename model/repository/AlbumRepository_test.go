package repository_test

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	e "web/service/gin/model/entities"
	mocks "web/service/gin/model/repository/mocks"
)

var expectedResult []e.Album = []e.Album{{
	ID:     uuid.MustParse("3fc7046e-666b-46a5-8028-b54f122118cf"),
	Title:  "Test_1",
	Artist: "TestA_1",
	Price:  10.0,
}, {
	ID:     uuid.MustParse("a362220d-0fdb-497c-8792-3aa0991f00fd"),
	Title:  "Test_2",
	Artist: "TestA_2",
	Price:  20.0,
}}

func TestRepository_GetAllAlbums(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAlbumRepo := mocks.NewMockAlbumRepositoryI(mockCtrl)
	mockAlbumRepo.EXPECT().GetAllAlbums().Return(expectedResult)

	result := mockAlbumRepo.GetAllAlbums()
	assert.Equal(t, expectedResult, result)
}
