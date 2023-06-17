package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceyMock struct {
	mock.Mock
}

func (s *serviceyMock) Create(newCompaignDto contract.NewCompaignDto) (string, error) {
	args := s.Called(newCompaignDto)
	return args.String(0), args.Error(1)
}

func (s *serviceyMock) GetBy(id string) (*contract.NewCompaignResponseDto, error) {
	//args := s.Called(id)
	return nil, nil
}

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCompaignDto{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@example.com"},
	}
	service := new(serviceyMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCompaignDto) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("12345", nil)
	handler := Handler{service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCompaignDto{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@example.com"},
	}
	service := new(serviceyMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
