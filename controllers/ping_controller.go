package controllers

import (
	"github.com/gin-gonic/gin"
)

//Ping Controller
type PingController struct{}

//new PingController
func NewPingController() *PingController {
	return &PingController{}
}

func (controller *PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}