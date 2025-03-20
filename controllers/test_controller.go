package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

// NewTestController creates a new instance of TestController
func NewTestController() *TestController {
	return &TestController{}
}

// Ping handles a simple ping-pong response
//
//	@Summary		Ping Test
//	@Description	Simple ping-pong endpoint for testing
//	@Tags			Test
//	@Produce		json
//	@Success		200	{string}	string	"pong"
//	@Router			/test/ping [get]
func (ctr *TestController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
