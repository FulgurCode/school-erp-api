package teacherHelpers

import (
	"strconv"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Checking if teacher user exists
func UserExists(teacher map[string]interface{}) bool {
	_, used := teacher["password"].(string)
	return !used
}

// Set up signup otp
func SignUpSetOTP(c *gin.Context, data map[string]interface{}) {
	// Creating OTP and storing in session
	var otp = helpers.CreateOTP()
	var session = sessions.DefaultMany(c, "teacherSignupOTP")
	session.Set("otp", otp)
	session.Options(sessions.Options{MaxAge: 120})
	// Storing teacher details in session
	session.Set("teacherEmail", data["email"].(string))
	session.Set("teacherPassword", data["password"].(string))
	session.Save()
	// Sending OTP to email and Response for request
	helpers.SendOTP(otp, data["email"].(string))
}

// Compare OTP
func CompareOtp(c *gin.Context) bool {
	var otp, _ = strconv.Atoi(c.Query("otp"))
	var session = sessions.DefaultMany(c, "teacherSignupOTP")
	if otp == session.Get("otp") {
		return true
	}
	return false
}

// Get data and Create teacher user
func CreateTeacherUser(c *gin.Context) error {
	var session = sessions.DefaultMany(c, "teacherSignupOTP")
	var email = session.Get("teacherEmail")
	var password = session.Get("teacherPassword")
	password = helpers.HashPassword(password.(string))
	var teacher = map[string]interface{}{"email": email, "password": password}
	var err = databaseHelpers.UpdateTeacherWithMail(bson.M{"email": email}, teacher)
	return err
}
