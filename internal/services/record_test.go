package services

import (
	"errors"
	"net/http"
	"testing"

	"zincsearchstash/tests/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDataInsertRecord(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockZincSearch := mock.NewMockZincSearchService(controller)

	status := http.StatusOK

	mockZincSearch.EXPECT().Insert(gomock.Any()).Return(status, nil)

	service := RecordService{
		zincSearch: mockZincSearch,
	}

	msg := `{
        "City": "Turin"
    }`

	err := service.DataInsert(msg)

	assert.Nil(t, err)
}

func TestDataInsertRecordDataNilError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockZincSearch := mock.NewMockZincSearchService(controller)

	errorMock := errors.New("error, message arrived null")

	service := RecordService{
		zincSearch: mockZincSearch,
	}

	err := service.DataInsert(nil)

	assert.Equal(t, errorMock.Error(), err.Error())
}

func TestDataInsertRecordMarshalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockZincSearch := mock.NewMockZincSearchService(controller)

	errorMock := errors.New("json: unsupported type: chan int")
	dataMock := make(chan int)
	service := RecordService{
		zincSearch: mockZincSearch,
	}

	err := service.DataInsert(dataMock)

	assert.Equal(t, errorMock.Error(), err.Error())
}

func TestDataInsertRecordStatusError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockZincSearch := mock.NewMockZincSearchService(controller)

	errorMock := errors.New("error on insert data in zincsearch")

	mockZincSearch.EXPECT().Insert(gomock.Any()).Return(0, errorMock)

	service := RecordService{
		zincSearch: mockZincSearch,
	}

	msg := `{
        "City": "Turin"
    }`

	err := service.DataInsert(msg)

	assert.Equal(t, errorMock.Error(), err.Error())
}
