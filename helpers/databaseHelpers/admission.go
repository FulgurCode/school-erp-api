package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"github.com/FulgurCode/school-erp-api/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Getting admission number of last student
func GetLastAdmissionNumber() int32 {
	// database
	var db = connections.Db
	// Getting student with highest admission number
	var option = options.FindOne().SetSort(bson.M{"admissionNo": -1}).SetProjection(bson.M{"admissionNo": 1})
	var student map[string]interface{}
	var err = db.Collection("students").FindOne(context.Background(), bson.M{}, option).Decode(&student)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0
		}
		helpers.CheckNilErr(err)
	}
	return student["admissionNo"].(int32)
}

// Inserting student to database
func InsertStudent(student map[string]interface{}) error {
	// database
	var db = connections.Db
	// inserting students
	var _, err = db.Collection("students").InsertOne(context.Background(), student)
	return err
}

// Importing multiple students to database
func ImportStudents(students []interface{}) error {
	// database
	var db = connections.Db
	// Importing students
	var _, err = db.Collection("students").InsertMany(context.Background(), students)
	return err
}
