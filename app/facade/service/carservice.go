package service

import (
	"boring-project/facade/model"
)

type allCars []model.Car

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

func Test() []model.Car {
	return cars
}
