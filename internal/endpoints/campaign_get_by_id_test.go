package endpoints

import (
	"emailn/internal/contract"
	internalMock "emailn/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.NewCompaignResponseDto{
		ID:      "123",
		Name:    "teste",
		Content: "Hi everyone",
		Status:  "Pending",
	}
	service := new(internalMock.CompaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaign, nil)
	handler := Handler{service}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contract.NewCompaignResponseDto).ID)
	assert.Equal(campaign.Name, response.(*contract.NewCompaignResponseDto).Name)
}

func Test_CampaignGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalMock.CompaignServiceMock)
	errExpected := errors.New("Something wrong happened")
	service.On("GetById", mock.Anything).Return(nil, errExpected)
	handler := Handler{service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(rr, req)

	assert.Equal(errExpected.Error(), errReturned.Error())
}
