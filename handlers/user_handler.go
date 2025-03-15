package handlers

import (
	"net/http"
	"vocalborn/backend-go/models"

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

func Register(c *gin.Context) {
	var registerRequest models.AccountRegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
	
}
// Logout handles user logout
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout endpoint",
	})
}