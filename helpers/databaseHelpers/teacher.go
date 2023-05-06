package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Getting teacher details using teacher id
func GetTeacher(teacherId primitive.ObjectID) (map[string]interface{}, error) {
	// database
	var db = connections.Db
	// get teacher and return
	var teacher map[string]interface{}
	var err = db.Collection("teachers").FindOne(context.Background(), bson.M{"_id": teacherId}).Decode(&teacher)
	return teacher, err
}
