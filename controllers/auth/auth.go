package auth

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	
	"github.com/eriwu052/auth-system1/models"
	"github.com/eriwu052/auth-system1/utils/token"
)

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func loginCheck(email, password string) (string, error) {

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", nil
	}
	
	err = verifyPassword(user.PasswordHash, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := loginCheck(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

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

	_, err := u.SaveUser(input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
