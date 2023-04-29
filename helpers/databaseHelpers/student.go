package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
)

// Getting students using admission number
func GetStudentByAdmissionNo(query map[string]interface{}) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting students from database
	var result, err = db.Collection("students").Find(context.Background(), query)
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Getting students using admission number
func GetStudentByApplicationNo(query map[string]interface{}) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting student from database
	var result, err = db.Collection("students").Find(context.Background(), query)
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Getting students using name
func GetStudentByName(query map[string]interface{}) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting student from database
	var result, err = db.Collection("students").Find(context.Background(), query)
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}
