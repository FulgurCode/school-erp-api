package controller

import (
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/adminHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/studentHelpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// POST request on /api/admin/login
func AdminLoginRoute(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Checking for username
	var admin, err = databaseHelpers.GetAdminWithUsername(data["username"].(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(401, "Wrong username or password")
			return
		}
		c.JSON(500, "Request failed")
	}
	// Comparing password and sending response
	var result = helpers.ComparePassword(admin["password"].(string), data["password"].(string))
	if !result {
		c.JSON(401, "Wrong username or password")
		return
	}
	// storing id and sending response if password is correct
	adminHelpers.LoginWithSesssion(c, admin)
	c.JSON(200, "Login Successful")
}

// GET request on '/api/admin/checklogin'
func AdminCheckLoginRoute(c *gin.Context) {
	// checking if logged in as admin and sending response
	var isLoggedIn = adminHelpers.CheckLogin(c)
	c.JSON(200, isLoggedIn)
}

// DELETE request on '/api/admin/logout'
func AdminLogoutRoute(c *gin.Context) {
	// clearing 'admin' session
	adminHelpers.Logout(c)
	// Response for the request
	c.JSON(200, "Loggged Out")
}

// PUT request on '/api/admin/change-password'
func ChangeAdminPassword(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Changing admin password and sending response
	data["id"] = adminHelpers.GetId(c)
	var err = adminHelpers.ChangePassword(data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Password Changed")
}

// POST request on '/api/admin/new-admission'
func NewAdmissionRoute(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Adding student to database and sending response
	var err = adminHelpers.AddNewAdmission(data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Successfully added")
}

// GET request on '/api/admin/get-students'
func GetStudentsRoute(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting search details
	var search = c.Query("search")
	var value = c.Query("value")
	// Getting students and sending response
	var students, err = adminHelpers.GetStudents(search, value)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, students)
}

// POST request on '/api/admin/import-students'
func ImportStudents(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting uploaded data file
	var file, err = c.FormFile("file")
	helpers.CheckNilErr(err)
	var students = studentHelpers.ImportStudentsFromCSV(file)
	// importing students to database sending response
	err = databaseHelpers.ImportStudents(students)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Successfully added")
}

// PUT request on '/api/admin/edit-student'
func EditStudent(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting object id of student
	var studentId, err = primitive.ObjectIDFromHex(c.Query("studentId"))
	helpers.CheckNilErr(err)
	// Getting request body
	var data = helpers.GetRequestBody(c)
	err = adminHelpers.EditStudent(studentId, data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Updated Successfully")
}
