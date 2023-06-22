package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/internalerrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	internalMock "emailn/internal/test/internal-mock"
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalMock.CompaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := campaign.ServiceImpl{repositoryMock}
	newCampaign := contract.NewCompaignDto{
		Name:     "Test Y",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	service := campaign.ServiceImpl{}
	newCampaign := contract.NewCompaignDto{
		Name:     "",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}

	_, err := service.Create(newCampaign)

	assert.False(errors.Is(internalerrors.ErrInternalError, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:     "Test Y",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}
	repositoryMock := new(internalMock.CompaignRepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := campaign.ServiceImpl{repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:     "Test Y",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}
	repositoryMock := new(internalMock.CompaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save database"))

	service := campaign.ServiceImpl{repositoryMock}

	_, err := service.Create(newCampaign)

	assert.True(t, errors.Is(internalerrors.ErrInternalError, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:     "Test Y",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}

	camp, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreateBy)

	repositoryMock := new(internalMock.CompaignRepositoryMock)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == camp.ID
	})).Return(camp, nil)

	service := campaign.ServiceImpl{repositoryMock}

	campaignReturned, _ := service.GetById(camp.ID)

	assert.Equal(t, camp.ID, campaignReturned.ID)
	assert.Equal(t, camp.Name, campaignReturned.Name)
	assert.Equal(t, camp.Content, campaignReturned.Content)
	assert.Equal(t, camp.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:     "Test Y",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}

	camp, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreateBy)

	repositoryMock := new(internalMock.CompaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("Something wrong"))

	service := campaign.ServiceImpl{repositoryMock}

	_, err := service.GetById(camp.ID)

	assert.Equal(t, internalerrors.ErrInternalError.Error(), err.Error())
}

func Test_Delete_ReturnRecordNotFount_when_campaign_does_not_exist(t *testing.T) {
	newCampaign := contract.NewCompaignDto{
		Name:     "Test Y",
		Content:  "Body hi",
		Emails:   []string{"test@example.com"},
		CreateBy: "test@example.com",
	}

	camp, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreateBy)

	repositoryMock := new(internalMock.CompaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	service := campaign.ServiceImpl{repositoryMock}

	err := service.Delete(camp.ID)

	assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
}
