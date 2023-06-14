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
	var err = db.Collection("admin").FindOne(context.Background(), bson.M{"username": username}).Decode(&admin)
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

// Add duty to teacher
func AddDuty(data map[string]interface{}) error {
	// database
	var db = connections.Db
	var teacher, _ = primitive.ObjectIDFromHex(data["teacherId"].(string))
	var duty = data["duty"].(string)
	var _, err = db.Collection("duties").InsertOne(context.Background(), bson.M{"teacherId": teacher, "duty": duty})
	return err
}

// Get all duties
func GetDuties() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Get duties of all teachers
	var result, err = db.Collection("duties").Aggregate(context.Background(), []bson.M{
		{
			"$lookup": bson.M{
				"from":         "teachers",
				"localField":   "teacherId",
				"foreignField": "_id",
				"as":           "teacher",
			},
		},
		{
			"$unwind": "$teacher",
		},
		{
			"$project": bson.M{
				"teacher._id":   1,
				"teacher.name":  1,
				"teacher.penNo": 1,
				"duty":          1,
			},
		},
	})
	var duties []map[string]interface{}
	for result.Next(context.Background()) {
		var duty map[string]interface{}
		result.Decode(&duty)
		duties = append(duties, duty)
	}
	return duties, err
}

// Delete duty
func DeleteDuty(id primitive.ObjectID) error {
	// database
	var db = connections.Db
	var _, err = db.Collection("duties").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

// Get admin user with id
func GetAdmin(id primitive.ObjectID) (map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Getting admin data
	var admin map[string]interface{}
	var err = db.Collection("admin").FindOne(context.Background(), bson.M{"_id": id}).Decode(&admin)
	return admin, err
}
