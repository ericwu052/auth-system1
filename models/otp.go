package models

import (
	"strings"
	"math/rand"
	
	"gorm.io/gorm"

	"github.com/eriwu052/auth-system1/utils/output"
)

type Otp struct {
	gorm.Model
	/* reset, checkMail, checkPhone */
	OtpType string `gorm:"size:255;not null" json:"otp_type"`
	OtpNumber string `gorm:"size:255;not null" json:"otp_number"`
	UserId uint
}

func generateRandomOtp(length int) string {
	charSet := "0123456789"
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

/** returns true when successfully set forgot password OTP */
func SetNewForgotPasswordOtp(user_id uint) (string, error) {
	mOtp := Otp{
		OtpType: "reset",
		OtpNumber: generateRandomOtp(6),
		UserId: user_id,
	}
	err := GlobalDb.Create(&mOtp).Error
	if err != nil {
		// failed to set OTP
		return "", err
	}

	return mOtp.OtpNumber, nil
}

func writeOtpToOutput(otpNumber string) {
	output.OutputOtp(otpNumber)
}

func ForgotPasswordFlow(user_id uint) (bool, error) {
	otp, err := SetNewForgotPasswordOtp(user_id)
	if err != nil {
		return false, err
	}

	writeOtpToOutput(otp)
	/** for email/mobile phone OTP, we should call sendMail / sendSMS too */
	return true, nil
}

func GetLatestOtp(user_id uint) (string, error) {
	var lastOtp Otp
	err := GlobalDb.Model(Otp{}).
		Where("user_id = ?", user_id).
		Last(&lastOtp).
		Error
	if err != nil {
		// otp not found for user
		return "", err
	}

	return lastOtp.OtpNumber, nil
}
