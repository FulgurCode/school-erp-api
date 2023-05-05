package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Update admin password
func UpdateAdminPassword(id primitive.ObjectID, password string) error {
	// database
	var db = connections.Db
	// update password and return error
	var _, err = db.Collection("admin").UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"password": password}})
	return err
}

// Add teacher to database
func AddTeacher(teacher map[string]interface{}) error {
	// database
	var db = connections.Db
	// Insert teacher
	var _, err = db.Collection("teachers").InsertOne(context.Background(), teacher)
	return err
}
