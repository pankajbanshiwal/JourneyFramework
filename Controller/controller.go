package Controller

import (
	"JourneyFramework/Service"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func Init() {
	var operationType int
	fmt.Printf("\n******Welcome to Journey Framework******\n\n")
	fmt.Printf("Starting Journey framework Inputs -> \n 1: createJourney (Again enter number of journeys to be created) \n 2: updateState \n 3: getJourney \n 4: evaluate \n 5: getCurrentStage \n 6: isOnboarded \n 10: terminating\n\n")

	breakLoop := false
	for {
		if breakLoop {
			break
		}
		fmt.Println("Enter value for next opration: ")
		fmt.Scan(&operationType)
		//ManageLibraryOperations(operationType)

		switch operationType {
		case 1:
			{
				//fmt.Println("Creating journey: ", operationType)
				var numOfJourneys int
				fmt.Println("Enter num of journeys to be creatd: ")
				fmt.Scan(&numOfJourneys)
				//CreateJouney(numberOfStages int, id, name, state, startDate, endDate string) bool {

				for i := 1; i <= numOfJourneys; i++ {
					id := uuid.New()
					name := "journey" + cast.ToString(i)
					Service.CreateJouney(5, id.String(), name, "CREATED", "", "")
					// error handlign here
				}
				fmt.Println("Journeys created = ", Service.Journies)

				//resp := Service.GetJourney(Service.Journies[3].ID) // id to be pass here
				//fmt.Println("Fetched journey = ", resp)

				// updating
				//Service.UpdateState(Service.Journies[3].ID, "TERMINATED")
				//fmt.Println("Total journeys = ", Service.Journies)

				// evaluate
				/*
					id := uuid.New()
					payload := make(map[string]string)
					payload["name"] = "pankaj"
					payload["city"] = "bangalore"
					payload["age"] = ">25"

					Service.Evaluate(id.String(), payload)
					fmt.Println("User journeys = ", Service.UserJourneys)

					payload = make(map[string]string)
					payload["name"] = "pankaj"
					//payload["city"] = "bangalore"
					payload["age"] = ">25"

					for i := 0; i < 4; i++ {
						Service.Evaluate(id.String(), payload)
						fmt.Println("User journeys = ", Service.UserJourneys)
					}
				*/

			}
		case 2:
			{
				id := ""
				fmt.Println("Enter journey id: ")
				fmt.Scan(&id)
				//fmt.Println("Update journey: ", operationType)
				Service.UpdateState(id, "TERMINATED")
				//handle error code
				fmt.Println("Current journies = ", Service.Journies)
			}
		case 3:
			{
				//fmt.Println("get journey: ", operationType)
				//GetJourney
				resp := Service.GetJourney("id") // id to be pass here
				fmt.Println("Fetched journey = ", resp)

			}
		case 4:
			{
				//fmt.Println("Enter value for next opration: ", operationType)
				Service.Evaluate("id", nil)
				fmt.Println("User journeys = ", Service.UserJourneys)

			}
		case 5:
			{
				fmt.Println("Enter value for next opration: ", operationType)

			}
		case 6:
			{
				fmt.Println("Enter value for next opration: ", operationType)

			}
		case 10:
			{
				fmt.Println("terminating: ", operationType)
				breakLoop = true

			}
		}
	}
	fmt.Println("Module has been terminated successfully!!")

}
