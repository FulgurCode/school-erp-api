package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"github.com/FulgurCode/school-erp-api/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var admissionNo, exists = student["admissionNo"].(int32)
	if exists == false {
		admissionNo = 0
	}
	return admissionNo
}

// Inserting student to database
func InsertStudent(student map[string]interface{}) (string, error) {
	// database
	var db = connections.Db
	// inserting students
	var result, err = db.Collection("students").InsertOne(context.Background(), student)
	var id = result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

// Importing multiple students to database
func ImportStudents(students []interface{}) error {
	// database
	var db = connections.Db
	// Importing students
	var _, err = db.Collection("students").InsertMany(context.Background(), students)
	return err
}

// Update student details using object id
func UpdateStudent(admissionId primitive.ObjectID, student map[string]interface{}) error {
	// database
	var db = connections.Db

	// Updating student details based on student id
	var _, err = db.Collection("students").UpdateOne(context.Background(), bson.M{"_id": admissionId}, bson.M{"$set": student})
	helpers.CheckNilErr(err)
	return err
}

// Importing multiple teachers to database
func ImportTeachers(teachers []interface{}) error {
	// database
	var db = connections.Db
	// Importing teachers
	var _, err = db.Collection("teachers").InsertMany(context.Background(), teachers)
	return err
}

// Get course and language report
func CourseLanguageReport() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Gettnig data from database
	var result, err = db.Collection("students").Aggregate(context.Background(), []bson.M{
		{
			"$group": bson.M{
				"_id":   bson.M{"course": "$course", "secondLanguage": "$secondLanguage"},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"course":         "$_id.course",
				"count":          1,
				"secondLanguage": "$_id.secondLanguage",
				"_id":            0,
			},
		},
	})
	var datas []map[string]interface{}
	for result.Next(context.Background()) {
		var data map[string]interface{}
		result.Decode(&data)
		datas = append(datas, data)
	}
	return datas, err
}

// Get course and status report
func CourseStatusReport() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	var result, err = db.Collection("students").Aggregate(context.Background(), []bson.M{
		{
			"$group": bson.M{
				"_id":   bson.M{"course": "$course", "status": "$status"},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"course": "$_id.course",
				"count":  1,
				"status": "$_id.status",
				"_id":    0,
			},
		},
	})
	var datas []map[string]interface{}
	for result.Next(context.Background()) {
		var data map[string]interface{}
		result.Decode(&data)
		datas = append(datas, data)
	}
	return datas, err
}

// Get course and gender report
func CourseGenderReport() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	var result, err = db.Collection("students").Aggregate(context.Background(), []bson.M{
		{
			"$group": bson.M{
				"_id":   bson.M{"course": "$course", "gender": "$gender"},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"course": "$_id.course",
				"count":  1,
				"gender": "$_id.gender",
				"_id":    0,
			},
		},
	})
	var datas []map[string]interface{}
	for result.Next(context.Background()) {
		var data map[string]interface{}
		result.Decode(&data)
		datas = append(datas, data)
	}
	return datas, err
}
