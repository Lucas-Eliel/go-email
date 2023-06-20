package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternalError = errors.New("Internal server error")

func ProcessErrorToReturn(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternalError
	}
	return err
}
