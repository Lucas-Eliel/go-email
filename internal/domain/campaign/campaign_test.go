package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaswdr/faker"
)

// O test de unidade tem os 3 A´s -> Arrange, Act, Assert
func Test_NewCampaign(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body Hi"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content)
	assert.Equal(t, len(contacts), len(campaign.Contacts))
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	//Arrange
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	_, err := NewCampaign("", content, contacts)

	//Assert
	assert.Equal(t, "name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	//Arrange
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	fake := faker.New()

	//Act
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	//Assert
	assert.Equal(t, "name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	//Arrange
	name := "Campaign X"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	_, err := NewCampaign(name, "", contacts)

	//Assert
	assert.Equal(t, "content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	//Arrange
	name := "Campaign X"
	contacts := []string{"email1@e.com", "email2@e.com"}
	fake := faker.New()

	//Act
	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	//Assert
	assert.Equal(t, "content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body hi"

	//Act
	_, err := NewCampaign(name, content, nil)

	//Assert
	assert.Equal(t, "contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body hi"

	//Act
	_, err := NewCampaign(name, content, []string{"email_invalid"})

	//Assert
	assert.Equal(t, "email is invalid", err.Error())
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body Hi"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body hi"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.CreatedOn)
}

func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {
	//Arrange
	name := "Campaign X"
	content := "Body hi"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(t, Pending, campaign.Status)
}
