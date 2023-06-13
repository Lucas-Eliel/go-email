package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// O test de unidade tem os 3 A´s -> Arrange, Act, Assert
func TestNewCampaign(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(t, "1", campaign.ID)
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content)
	assert.Equal(t, len(contacts), len(campaign.Contacts))
}

func TestNewCampaignIDIsNotNil(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.ID)
}

func TestNewCampaignCreatedOnIsNotNil(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.CreatedOn)
}
