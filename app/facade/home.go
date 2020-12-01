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

  v1 := router.Group("/v1")
  {
    v1.GET("/events", getAllEvents)
  }

  v2 := router.Group("/v2")
  {
    v2.GET("/internal", getAllEvents)
  }
	router.Run(":8090")
}

func main() {
	handleRequest()
}
