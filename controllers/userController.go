package controllers

import (
	"net/http"

	"github.com/chandchand/be-hms-go/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Profile(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := services.GetUserByID(uuid.MustParse(userID.(string)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed retrieving user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true", "message": "success","data": user})
}