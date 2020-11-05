package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"boring-project/facade/service"
	"boring-project/facade/model"

)


type allEvents []model.Event

var events = allEvents{
	{
		ID:          "1",
		Title:       "introduction to golang",
		Description: "This is a test",
	},
}

func getAllEvents(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	// var temp=service.GetAllCars()

	var cars = service.Test()
	g.JSON(http.StatusOK, cars)
}

func handleRequest() {
	router := gin.Default()

	router.GET("/events", getAllEvents)

	router.Run(":8090")
}

func main() {
	handleRequest()
}
