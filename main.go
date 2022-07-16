package main

import (
	"github.com/andey-robins/orchestrator/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/alive", api.Alive)
	r.GET("/status", api.Status)
	r.Run()
}
