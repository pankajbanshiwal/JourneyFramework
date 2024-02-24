package Repositories

type Journey struct {
	ID             string // uuid
	Name           string
	NumberOfStages int // 3 stage
	CurrentStage   int
	StartDate      string
	EndDate        string
	State          string // CREATED , ACTIVE , TERMINATED , EXPIRED
}

type UserJournies struct {
	UserId         string // uuid
	ActiveJourneys []Journey
}
