package contract

type NewCompaignResponseDto struct {
	ID                   string
	Name                 string
	Content              string
	Status               string
	AmountOfEmailsToSend int
}
