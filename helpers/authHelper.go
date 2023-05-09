package helpers

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

// Comparing password and creating response
func ComparePassword(hashedPassword string, password string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

// Hashing password using bcrypt
func HashPassword(password string) string {
	// Generating hashed password
	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

// Create OTP
func CreateOTP() int {
	// Generate a random number between 100000 and 999999
	rand.Seed(time.Now().UnixNano())
	var otp = rand.Intn(900000) + 100000
	return otp
}

// send OTP to email address
func SendOTP(otp int, email string) error {
	// Getting mail and password (should be zohomail)
	var sendMailAddress = os.Getenv("MAIL_ADDRESS")
	var sendMailPassword = os.Getenv("MAIL_PASSWORD")
	// Creating email
	var m = gomail.NewMessage()
	m.SetHeader("From", sendMailAddress)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "CHSS Authentication OTP")
	var body = fmt.Sprintf("<h3>This is the OTP for your CHSS chattanchal website account creation </h3><br><h1> %d </h1>", otp)
	m.SetBody("text/html", body)
	// Sendig email
	var dialer = gomail.NewDialer("smtp.gmail.com", 465, sendMailAddress, sendMailPassword)
	var err = dialer.DialAndSend(m)
	return err
}

// Compare OTP
func CompareOTP(enteredOTP string, session sessions.Session) bool {
	var otpInt, _ = strconv.Atoi(enteredOTP)
	var otp = session.Get("otp").(int)
	if otp == otpInt {
		return true
	} else {
		return false
	}
}

// Creating response for otp if checking failed
func OtpFailedResponse(moreTry bool, reason string) map[string]interface{} {
	var response = map[string]interface{}{"status": false, "reason": reason, "moreTry": moreTry}
	return response
}
