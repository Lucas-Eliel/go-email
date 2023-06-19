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

func (r *repositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
}

func (r *repositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := ServiceImpl{repositoryMock}
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

	service := ServiceImpl{}
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

	service := ServiceImpl{repositoryMock}

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

	service := ServiceImpl{repositoryMock}

	_, err := service.Create(newCampaign)

	assert.True(t, errors.Is(internalerrors.ErrInternalError, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:    "Test Y",
		Content: "Body hi",
		Emails:  []string{"test@example.com"},
	}

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)

	service := ServiceImpl{repositoryMock}

	campaignReturned, _ := service.GetById(campaign.ID)

	assert.Equal(t, campaign.ID, campaignReturned.ID)
	assert.Equal(t, campaign.Name, campaignReturned.Name)
	assert.Equal(t, campaign.Content, campaignReturned.Content)
	assert.Equal(t, campaign.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:    "Test Y",
		Content: "Body hi",
		Emails:  []string{"test@example.com"},
	}

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("Something wrong"))

	service := ServiceImpl{repositoryMock}

	_, err := service.GetById(campaign.ID)

	assert.Equal(t, internalerrors.ErrInternalError.Error(), err.Error())
}
