package main

import (
	"net/http"

	"boring-project/facade/service"

	"github.com/gin-gonic/gin"
)

func getAllEvents(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	var cars = service.GetAllCars()
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
