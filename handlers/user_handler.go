package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// User Login Handler godoc
//	@Summary		User Login
//	@Description	User Login
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal server error"
//	@Router			/user/login [get]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login endpoint",
	})
}

// Register Handles GODOC
//	@Summary		User Register
//	@Description	User Register
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal server error"
//	@Router			/user/register [get]
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register endpoint",
	})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout endpoint",
	})
}