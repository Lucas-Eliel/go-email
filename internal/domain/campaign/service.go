package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
)

type Service interface {
	Create(newCompaignDto contract.NewCompaignDto) (string, error)
	GetById(id string) (*contract.NewCompaignResponseDto, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCompaignDto contract.NewCompaignDto) (string, error) {

	compaign, err := NewCampaign(newCompaignDto.Name, newCompaignDto.Content, newCompaignDto.Emails)

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
		return nil, internalerrors.ErrInternalError
	}

	return &contract.NewCompaignResponseDto{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
