package main

import (
	"net/http"

	"boring-project/event/service"

	"github.com/gin-gonic/gin"
)

func getAllEvents(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	var events = service.GetAllEvents()
	g.JSON(http.StatusOK, events)
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
