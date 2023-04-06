package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
)

// Get admin user from database
func GetAdminWithUsername(username string) (map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Get admin user and return
	var admin map[string]interface{}
	var result = db.Collection("admin").FindOne(context.Background(), bson.M{"username": username})
	var err = result.Err()
	result.Decode(&admin)
	return admin, err
}
