package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
)

// Getting students using admission number
func GetStudentByAdmissionNo(admissionNo int) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting students from database
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"admissionNo": admissionNo})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Getting students using admission number
func GetStudentByApplicationNo(applicationNo int) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting student from database
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"applicationNo": applicationNo})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Getting students using name
func GetStudentByName(name string) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting student from database
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"name": name})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}
