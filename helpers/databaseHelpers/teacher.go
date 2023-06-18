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

// Get all teachers
func GetAllTeachers() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	var result, err = db.Collection("teachers").Find(context.Background(), bson.M{})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Get duty
func GetDuty(id primitive.ObjectID, dutyName string) (map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Get teacher duty
	var duty map[string]interface{}
	var err = db.Collection("duties").FindOne(context.Background(), bson.M{"teacherId": id, "duty": dutyName}).Decode(&duty)
	return duty, err
}

// Update teacher password
func UpdateTeacherPassword(id primitive.ObjectID, password string) error {
	// database
	var db = connections.Db
	// update password and return error
	var _, err = db.Collection("teachers").UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"password": password}})
	if err != nil {
		panic(err)
	}
	return err
}

// Get all teacher duties from database
func GetTeacherDuties(id primitive.ObjectID) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Get teacher duty
	var duties []map[string]interface{}
	// get all duties of the teacher
	var result, err = db.Collection("duties").Find(context.Background(), bson.M{"teacherId": id})
	for result.Next(context.Background()) {
		var duty map[string]interface{}
		result.Decode(&duty)
		duties = append(duties, duty)
	}
	return duties, err
}
