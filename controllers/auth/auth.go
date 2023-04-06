package auth

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/eriwu052/auth-system1/models"
)

type RegisterInput struct {
	Email string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	MobileNo string `json:"mobileNo" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Fullname = input.Fullname
	u.MobileNo = input.MobileNo
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
