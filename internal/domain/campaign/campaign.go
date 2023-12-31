package campaign

import (
	"emailn/internal/internalerrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	CreateBy  string    `validate:"required" gorm:"size:50"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
}

const (
	Pending  string = "Pending"
	Started  string = "Started"
	Done     string = "Done"
	Canceled string = "Canceled"
	Deleted  string = "Deleted"
)

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func NewCampaign(name string, content string, emails []string, createBy string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, v := range emails {
		contacts[i].ID = xid.New().String()
		contacts[i].Email = v
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		CreateBy:  createBy,
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
