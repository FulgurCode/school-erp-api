package adminHelpers

import (
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get objectId of admin user
func GetId(c *gin.Context) primitive.ObjectID {
	var session = sessions.DefaultMany(c, "admin")
	var id, err = primitive.ObjectIDFromHex(session.Get("id").(string))
	helpers.CheckNilErr(err)
	return id
}
