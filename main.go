package main

import (
	"log"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
    authController "github.com/eriwu052/auth-system1/controllers/auth"
	"github.com/eriwu052/auth-system1/models"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	models.ConnectDatabase()
	mGin := gin.Default();

	/** handle CORS because we're using 2 server from different host */
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, []string{"Authorization"}...)
	mGin.Use(cors.New(corsConfig))

	publicRoutes := mGin.Group("/api")
	publicRoutes.POST("/register", authController.Register)
	publicRoutes.POST("/login", authController.Login)
	publicRoutes.POST("/login-phone", authController.LoginPhone)
	publicRoutes.POST("/request-otp-email", authController.RequestOtpEmail)
	
	mGin.Run()
}
