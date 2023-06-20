package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CompaignRepositoryMock struct {
	mock.Mock
}

func (r *CompaignRepositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CompaignRepositoryMock) Get() ([]campaign.Campaign, error) {
	//args := r.Called()
	return nil, nil
}

func (r *CompaignRepositoryMock) GetById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), nil
}

func (r *CompaignRepositoryMock) Update(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CompaignRepositoryMock) Delete(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
