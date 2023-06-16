package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called()
	return nil, nil
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := Service{repositoryMock}
	newCampaign := contract.NewCompaignDto{
		Name:    "Test Y",
		Content: "Body hi",
		Emails:  []string{"test@example.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	service := Service{}
	newCampaign := contract.NewCompaignDto{
		Name:    "",
		Content: "Body hi",
		Emails:  []string{"test@example.com"},
	}

	_, err := service.Create(newCampaign)

	assert.False(errors.Is(internalerrors.ErrInternalError, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:    "Test Y",
		Content: "Body hi",
		Emails:  []string{"test@example.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:    "Test Y",
		Content: "Body hi",
		Emails:  []string{"test@example.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save database"))

	service := Service{repositoryMock}

	_, err := service.Create(newCampaign)

	assert.True(t, errors.Is(internalerrors.ErrInternalError, err))
}
