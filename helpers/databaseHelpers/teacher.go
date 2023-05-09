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

// Getting teacher using email
func GetTeacherWithEmail(email string) (map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Get teacher from database
	var teacher map[string]interface{}
	var err = db.Collection("teachers").FindOne(context.Background(), bson.M{"email": email}).Decode(&teacher)
	return teacher, err
}

// Create Teacher user
func CreateTeacher(teacher map[string]interface{}) error {
	// database
	var db = connections.Db
	var _, err = db.Collection("teachers").InsertOne(context.Background(), teacher)
	return err
}

// Create Teacher user
func UpdateTeacherWithMail(search map[string]interface{}, teacher map[string]interface{}) error {
	// database
	var db = connections.Db
	var _, err = db.Collection("teachers").UpdateOne(context.Background(), search, bson.M{"$set": teacher})
	return err
}
