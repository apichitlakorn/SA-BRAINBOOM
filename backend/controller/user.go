package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Tawunchai/review-project/config"
	"github.com/Tawunchai/review-project/entity"
)

func GetUserById(c *gin.Context) {
	ID := c.Param("id")
	var user entity.User

	db := config.DB()
	results := db.Preload("UserRole").First(&user, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}

