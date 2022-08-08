package useCase

import (
	entity "travel_organizer_backend/app/domain/entities"
)

func GenerateTravel() entity.TravelEntity {
	var travel = entity.TravelEntity{
		Country: "Portugal",
		City:    "Lisbon",
	}
	return travel
}
