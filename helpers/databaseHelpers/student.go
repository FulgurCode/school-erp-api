package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
)

// Getting student details from database usign query
func GetStudents(query map[string]interface{}) ([]map[string]interface{}, error) {
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
