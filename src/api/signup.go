package api

import (
	"login-crud/config"
	"login-crud/database"
	"login-crud/src/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup function is used to create the account in the database
func SignUp(c *gin.Context) {
	var userdata database.Users
	err := c.Bind(&userdata)
	if err != nil {
		config.Error(err)
		return
	}
	password := userdata.Password
	hash, _ := auth.HashPassword(password)
	confirmpassword := userdata.ConfirmPassword
	status := auth.CheckPasswordHash(confirmpassword, hash)
	if !status {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "incorrect pass",
		})
		return
	}
	data := database.Users{
		Name:            userdata.Name,
		Email:           userdata.Email,
		Password:        hash,
		ConfirmPassword: hash,
	}
	result := config.DB.Create(&data)
	if result.Error != nil {
		config.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "account created successfully",
	})

}
