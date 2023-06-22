package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
	"errors"
)

type Service interface {
	Create(newCompaignDto contract.NewCompaignDto) (string, error)
	GetById(id string) (*contract.NewCompaignResponseDto, error)
	Delete(id string) error
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCompaignDto contract.NewCompaignDto) (string, error) {

	compaign, err := NewCampaign(newCompaignDto.Name, newCompaignDto.Content, newCompaignDto.Emails, newCompaignDto.CreateBy)

	if err != nil {
		return "", err
	}

	erro := s.Repository.Save(compaign)

	if erro != nil {
		return "", internalerrors.ErrInternalError
	}

	return compaign.ID, nil
}

func (s *ServiceImpl) GetById(id string) (*contract.NewCompaignResponseDto, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	return &contract.NewCompaignResponseDto{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *ServiceImpl) Delete(id string) error {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Delete()

	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerrors.ErrInternalError
	}

	return nil
}
