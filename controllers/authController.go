package controllers

import (
	"fmt"
	"net/http"

	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/chandchand/be-hms-go/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *gin.Context) {
    var user models.Staff
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Login.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate password hash", "error": err.Error()})
        return
    }

    user.Login.Password = string(hashedPass)

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create user", "error": err.Error()})
        return
    }

    fmt.Println("Stored Password:", user.Login.Password)

    c.JSON(http.StatusOK, gin.H{"success": true, "message": "User created successfully", "data": user})
}



func Login(c *gin.Context) {
	var loginData models.Login
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Staff
	database.DB.Preload("Login").Joins("JOIN logins ON staffs.id = logins.staff_id").Where("logins.username = ?", loginData.Username).First(&user)

	// fmt.Println(user.ID)
	// fmt.Println("Input Password:", loginData.Password)
	// fmt.Println("Stored Password:", user.Login.Password)
	// fmt.Println("Staff ID:", user.ID)
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Login.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID.String(), user.FirstName, user.Login.Username, user.Login.Role)
	c.JSON(http.StatusOK, gin.H{"token": token, "data": user})
}
