package models

import (
	"fmt"
	"log"
	"os"
	
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var GlobalDb *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbConstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	GlobalDb, err = gorm.Open(mysql.Open(dbConstr), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to MySQL database")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Connected to MySQL database")
	}

	GlobalDb.AutoMigrate(&User{})
	GlobalDb.AutoMigrate(&Otp{})
}
