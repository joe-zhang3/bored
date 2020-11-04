package service

import (
)

type car struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allCars []car

var cars = allCars{
	{
		ID:          "1",
		Title:       "cars introduction to golang",
		Description: "This is a car",
	},
}

func GetAllCars() allCars{
  return cars
}


