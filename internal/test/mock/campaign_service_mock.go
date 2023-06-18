package mock

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CompaignServiceMock struct {
	mock.Mock
}

func (s *CompaignServiceMock) Create(newCompaignDto contract.NewCompaignDto) (string, error) {
	args := s.Called(newCompaignDto)
	return args.String(0), args.Error(1)
}

func (s *CompaignServiceMock) GetById(id string) (*contract.NewCompaignResponseDto, error) {
	args := s.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.NewCompaignResponseDto), nil
}
