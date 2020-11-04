package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
  "boring-project/facade/service"
)

type event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []event

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
	var temp=service.GetAllCars()
	g.JSON(http.StatusOK, temp)
}

func handleRequest() {
	router := gin.Default()

	router.GET("/events", getAllEvents)

	router.Run(":8090")
}

func main() {
	handleRequest()
}
