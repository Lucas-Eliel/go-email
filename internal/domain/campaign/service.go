package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCompaignDto contract.NewCompaignDto) (string, error) {

	compaign, err := NewCampaign(newCompaignDto.Name, newCompaignDto.Content, newCompaignDto.Emails)

	if err != nil {
		return "", err
	}

	erro := s.Repository.Save(compaign)

	if erro != nil {
		return "", erro
	}

	return compaign.ID, nil
}
