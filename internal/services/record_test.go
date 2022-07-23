package services

import (
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
