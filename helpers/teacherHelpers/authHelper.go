package teacherHelpers

import (
	"strconv"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Storing teacher id in session
func LoginWithSesssion(c *gin.Context, teacher map[string]interface{}) {
	var session = sessions.DefaultMany(c, "teacher")
	session.Set("isLoggedIn", true)
	session.Set("id", teacher["_id"].(primitive.ObjectID).Hex())
	// Saving session for 1 years
	var maxAge = 60 * 60 * 24 * 365 * 100
	session.Options(sessions.Options{MaxAge: maxAge})
	var err = session.Save()
	helpers.CheckNilErr(err)
}

// Checking if teacher user exists
func UserExists(teacher map[string]interface{}) bool {
	_, used := teacher["password"].(string)
	return !used
}

// Set up signup otp
func SignUpSetOTP(c *gin.Context, data map[string]interface{}) error {
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
	var err = helpers.SendOTP(otp, data["email"].(string))
	return err
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

// Check if teacher logged in
func CheckLogin(c *gin.Context) bool {
	var session = sessions.DefaultMany(c, "teacher")
	if session.Get("isLoggedIn") == true {
		return true
	}
	return false
}

// Logging out as teacher using session
func Logout(c *gin.Context) {
	// clearing teacher session
	var session = sessions.DefaultMany(c, "teacher")
	session.Options(sessions.Options{MaxAge: -1})
	session.Clear()
	session.Save()
}

// Getting id
func GetId(c *gin.Context) primitive.ObjectID {
	var session = sessions.DefaultMany(c, "teacher")
	var id, _ = primitive.ObjectIDFromHex(session.Get("id").(string))
	return id
}

// Update teacher password
func ChangePassword(teacher map[string]interface{}) error {
	// Hashed new password
	var hashedPassword = helpers.HashPassword(teacher["new-password"].(string))
	// Updating password and returning error
	var err = databaseHelpers.UpdateTeacherPassword(teacher["id"].(primitive.ObjectID), hashedPassword)
	return err
}
