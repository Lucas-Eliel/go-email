package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// O test de unidade tem os 3 A´s -> Arrange, Act, Assert
func Test_NewCampaign(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content)
	assert.Equal(t, len(contacts), len(campaign.Contacts))
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	//Arrange
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	_, err := NewCampaign("", content, contacts)

	//Assert
	assert.Equal(t, "name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	//Arrange
	name := "Campaign X"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	_, err := NewCampaign(name, "", contacts)

	//Assert
	assert.Equal(t, "content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"

	//Act
	_, err := NewCampaign(name, content, []string{})

	//Assert
	assert.Equal(t, "contacts is required", err.Error())
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.CreatedOn)
}
