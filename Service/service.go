package Service

import (
	"JourneyFramework/Repositories"
	"errors"
	"fmt"
)

var Journies []Repositories.Journey
var UserJourneys []Repositories.UserJournies

//func NewJouney(numberOfStages int, id, state, startDate, endDate string) *Journey {

func CreateJouney(numberOfStages int, id, name, state, startDate, endDate string) Repositories.Journey {
	resp := Repositories.NewJouney(numberOfStages, id, name, state, startDate, endDate)
	Journies = append(Journies, resp)
	return resp
}

func GetJourney(id string) *Repositories.Journey {
	// journey from id
	for _, val := range Journies {
		if val.ID == id {
			return &val
		}
	}
	return nil
}

func GetJourneyByName(name string) *Repositories.Journey {
	// journey from id
	for _, val := range Journies {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

func GetUserJourneyByName(userId string, name string) *Repositories.Journey {
	// journey from id
	for _, val := range UserJourneys {
		if val.UserId == userId {
			for _, v := range val.ActiveJourneys {
				if v.Name == name {
					return &v
				}
			}
		}
	}
	return nil
}

func UpdateState(id, state string) (bool, error) {
	// update the joruney state
	for ind, val := range Journies {
		if val.ID == id {
			val.State = state
			Journies[ind] = val
			return true, nil
		}
	}
	return false, errors.New("Error updating journey")
}

func UpdateUserJourney(userId, name, state string, stage int) (bool, error) {
	// update the joruney state
	for _, val := range UserJourneys {
		// loop jounrneys
		if val.UserId == userId {
			for ind, v := range val.ActiveJourneys {
				if v.Name == name {
					fmt.Println("HERE ==== ")
					v.State = state
					v.CurrentStage = stage
					val.ActiveJourneys[ind] = v
				}
			}
		}
	}
	return false, errors.New("Error updating journey")
}

func Evaluate(userId string, payload map[string]string) {
	// if the location id bangalore onboard him
	if val, ok := payload["city"]; ok {
		// onboard him
		if val == "bangalore" {
			userJourney := Repositories.UserJournies{}
			//journey := Repositories.Journey{}
			// fetch journey to be onboarded
			journey := GetJourneyByName("journey1")
			journey.CurrentStage = 1
			journey.State = "ACTIVE"
			userJourney.UserId = userId
			userJourney.ActiveJourneys = append(userJourney.ActiveJourneys, *journey)
			UserJourneys = append(UserJourneys, userJourney)
		}
	}

	if val, ok := payload["age"]; ok {
		// onboard him
		if val == ">25" {
			// fetch journey to be onboarded
			journey := GetUserJourneyByName(userId, "journey1")
			if journey.CurrentStage == journey.NumberOfStages {
				journey.State = "TERMINATED"
			} else {
				journey.CurrentStage += 1
			}
			UpdateUserJourney(userId, "journey1", journey.State, journey.CurrentStage)
		}
		return
	}
	//// else do nothing

}
