package service

import (
	"boring-project/event/model"
)

type AllEvents []model.Event

var events = AllEvents{
	{
		ID:          "1",
		Title:       "cars introduction to golang",
		Description: "This is a car",
	},
}

func GetAllEvents() AllEvents {
	return events
}
