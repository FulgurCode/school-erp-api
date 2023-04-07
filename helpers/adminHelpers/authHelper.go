package adminHelpers

import (
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Storing admin id in session
func LoginWithSesssion(c *gin.Context, admin map[string]interface{}) {
	var session = sessions.DefaultMany(c, "admin")
	session.Set("isLoggedIn", true)
	session.Set("id", admin["_id"].(primitive.ObjectID).Hex())
	// Saving session for 1 years
	var maxAge = 60 * 60 * 24 * 365 * 100
	session.Options(sessions.Options{MaxAge: maxAge})
	var err = session.Save()
	helpers.CheckNilErr(err)
}

// Check if admin logged in
func CheckLogin(c *gin.Context) bool {
	var session = sessions.DefaultMany(c, "admin")
	if session.Get("isLoggedIn") == true {
		return true
	}
	return false
}

// Logging out as admin using session
func Logout(c *gin.Context) {
	// clearing admin session
	var session = sessions.DefaultMany(c, "admin")
	session.Options(sessions.Options{MaxAge: -1})
	session.Clear()
	session.Save()
}
