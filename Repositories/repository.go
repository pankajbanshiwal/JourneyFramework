package Repositories

// constructor
func NewJouney(numberOfStages int, id, name, state, startDate, endDate string) Journey {
	journey := Journey{}
	journey.NumberOfStages = numberOfStages
	journey.ID = id
	journey.Name = name
	journey.State = state
	journey.StartDate = startDate
	journey.EndDate = endDate
	return journey
}

// func UpdateState(id, state string) Repositories.Journey {

// }
