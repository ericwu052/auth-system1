package main

import "github.com/gin-gonic/gin"
import authController "github.com/eriwu052/auth-system1/controllers/auth"

func main() {
	mGin := gin.Default();
	mGin.POST("/register", authController.Register)
	mGin.Run()
}
