package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Getting students using name
func GetStudentsByName(name string, data map[string]interface{}) ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	var result, err = db.Collection("students").Aggregate(context.Background(), []bson.M{
		{
			"$match": bson.M{"$and": []bson.M{{"$text": bson.M{"$search": name}}, data}},
		},
	})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Getting student details using student id
func GetStudent(studentId primitive.ObjectID) (map[string]interface{}, error) {
	// database
	var db = connections.Db
	// get student and return
	var student map[string]interface{}
	var err = db.Collection("students").FindOne(context.Background(), bson.M{"_id": studentId}).Decode(&student)
	return student, err
}

// Get admitted students
func GetAdmittedStudents() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Getting admitted students details from database
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"status": bson.M{"$ne": "pending"}})
	var students []map[string]interface{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Get students remain to verify
func GetStudentsToVerify() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Getting admitted students details from database
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"status": bson.M{"$ne": "pending"}, "verified": bson.M{"$ne": true}})
	var students = []map[string]interface{}{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Verify student
func VerifyStudent(studentId primitive.ObjectID) error {
	// database
	var db = connections.Db
	// Verify student
	var _, err = db.Collection("students").UpdateOne(context.Background(), bson.M{"_id": studentId}, bson.M{"$set": bson.M{"verified": true}})
	return err
}

// Getting students that remain to confirm
func GetStudentsToConfirm() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	// Getting students that verified which is not confirmed
	var result, err = db.Collection("students").Find(context.Background(), bson.M{"status": bson.M{"$ne": "pending"}, "verified": true, "confirmed": bson.M{"$ne": true}})
	var students = []map[string]interface{}{}
	for result.Next(context.Background()) {
		var student map[string]interface{}
		result.Decode(&student)
		students = append(students, student)
	}
	return students, err
}

// Confirm student
func ConfirmStudent(studentId primitive.ObjectID) error {
	// database
	var db = connections.Db
	// Confirm student
	var _, err = db.Collection("students").UpdateOne(context.Background(), bson.M{"_id": studentId}, bson.M{"$set": bson.M{"confirmed": true}})
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
		{
			"$sort": bson.M{
				"course": -1,
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
		{
			"$sort": bson.M{
				"course": -1,
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
		{
			"$sort": bson.M{
				"course": -1,
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

// Get course and category report
func CourseCategoryReport() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	var result, err = db.Collection("students").Aggregate(context.Background(), []bson.M{
		{
			"$group": bson.M{
				"_id":   bson.M{"course": "$course", "category": "$category"},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"course":   "$_id.course",
				"count":    1,
				"category": "$_id.category",
				"_id":      0,
			},
		},
		{
			"$sort": bson.M{
				"course": -1,
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

// Get course and caste report
func CourseCasteReport() ([]map[string]interface{}, error) {
	// database
	var db = connections.Db
	var result, err = db.Collection("students").Aggregate(context.Background(), []bson.M{
		{
			"$group": bson.M{
				"_id":   bson.M{"course": "$course", "caste": "$caste"},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"course": "$_id.course",
				"count":  1,
				"caste":  "$_id.caste",
				"_id":    0,
			},
		},
		{
			"$sort": bson.M{
				"course": -1,
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
