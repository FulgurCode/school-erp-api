package controller

import (
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/teacherHelpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// POST request on '/api/teacher/signup'
func TeacherSignup(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	var teacher, err = databaseHelpers.GetTeacherWithEmail(data["email"].(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(401, "No teacher with this email")
			return
		}
		c.JSON(500, "Request failed")
		return
	}
	var exists = teacherHelpers.UserExists(teacher)
	if !exists {
		c.JSON(409, "Account already made")
	}
	err = teacherHelpers.SignUpSetOTP(c, data)
	if err != nil {
		c.JSON(500, "Network issue")
	}
	c.JSON(200, "OTP sended to the email adress")
}

// GET request on '/api/teacher/signup-otp'
func TeacherSignupOTP(c *gin.Context) {
	// comparing otp
	var result = teacherHelpers.CompareOtp(c)
	// Checking if password is correct and sending response
	if !result {
		c.JSON(401, "Incorrect OTP")
		return
	}
	teacherHelpers.CreateTeacherUser(c)
	c.JSON(200, "Teacher account created")
}

// POST request '/api/teacher/login'
func TeacherLogin(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Checking for username
	var teacher, err = databaseHelpers.GetTeacherWithEmail(data["email"].(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(401, "Wrong username or password")
			return
		}
		c.JSON(500, "Request failed")
	}
	// Comparing password and sending response
	var result = helpers.ComparePassword(teacher["password"].(string), data["password"].(string))
	if !result {
		c.JSON(401, "Wrong username or password")
		return
	}
	// storing id and sending response if password is correct
	teacherHelpers.LoginWithSesssion(c, teacher)
	c.JSON(200, "Login Successful")
}
