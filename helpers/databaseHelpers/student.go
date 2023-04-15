package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
)

// Getting student using admission number
func GetStudentByAdmissionNo(admissionNo int) (map[string]interface{}, error) {
	// database
	var db = connections.Db

	// Getting student from database
	var student map[string]interface{}
	var err = db.Collection("students").FindOne(context.Background(), bson.M{"admissionNo": admissionNo}).Decode(&student)
	return student, err
}
