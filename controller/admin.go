package controller

import (
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/adminHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/gin-gonic/gin"
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
	if isLoggedIn == true {
		c.JSON(200, true)
	} else {
		c.JSON(401, false)
	}
}

// GET request on '/api/admin/logout'
func AdminLogoutRoute(c *gin.Context) {
	// clearing 'admin' session
	adminHelpers.Logout(c)
	// Response for the request
	c.JSON(200, "Loggged Out")
}

// PUT request on '/api/admin/change-password'
func ChangeAdminPassword(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Changing admin password and sending response
	data["id"] = adminHelpers.GetId(c)
	var err = adminHelpers.ChangePassword(data)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "Password Changed")
}
