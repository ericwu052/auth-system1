package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Fullname string `gorm:"size:255;not null;" json:"fullname"`
	MobileNo string `gorm:"size:255;not null;" json:"mobileNo"`
	Password string
	PasswordHash string `gorm:"size:255;not null;" json:"passwordHash"`
}

func (this *User) SaveUser() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	this.PasswordHash = string(hashedPassword)

	err = GlobalDb.Create(&this).Error
	if err != nil {
		return &User{}, err
	}

	return this, nil
}
