package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllEvents(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	epid := g.GetHeader("X-EPID")

	if len(epid) == 0 {
		epid = "No epid passed in"
	}

	g.JSON(http.StatusOK, epid)
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
