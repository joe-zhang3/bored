package main

import (
  "net/http"
  "encoding/json"
  "github.com/gin-gonic/gin"
)

type event struct{
  ID  string `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
}

type allEvents []event

var events = allEvents{
  {
    ID: "1",
    Title: "introduction to golang",
    Description: "This is a test",
  },
}

func homePage(g *gin.Context){
  g.JSON(http.StatusOK, gin.H{"data": events})
}

func getAllEvents(w http.ResponseWriter, r *http.Request){
  json.NewEncoder(w).Encode(events)
}

func handleRequest(){
  router := gin.Default()

  router.GET("/", homePage)

  router.Run(":8090")
}

func main(){
  handleRequest()
}
