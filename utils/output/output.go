package output

import (
	"os"
)

var moduleFile *os.File

func initializeModuelFile() {
	moduleFile, _ = os.Create("output/output.txt")
}

func moduleFileUninitialized() bool {
	return true
}

func prepareModuleFile() {
	if moduleFileUninitialized() {
		initializeModuelFile()
	}
}

func OutputOtp(otp string) {
	prepareModuleFile()
	
	moduleFile.WriteString(otp)
	moduleFile.Sync()
}
