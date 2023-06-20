package databaseHelpers

import (
	"context"

	"github.com/FulgurCode/school-erp-api/connections"
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Getting admission number of last student
func GetLastAdmissionNumber() int {
	// database
	var db = connections.Db
	// Getting student with highest admission number
	var option = options.Find().SetSort(bson.M{"admissionNo": -1})
	var student models.Student
	var result, err = db.Collection("students").Find(context.Background(), bson.M{}, option)
	for result.Next(context.Background()) {
		result.Decode(&student)
		break
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0
		}
	}
	if student.AdmissionNo == 0 {
		return 0
	}
	return student.AdmissionNo
}

// Inserting student to database
func InsertStudent(student models.Student) (string, error) {
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
func UpdateStudent(admissionId primitive.ObjectID, student models.Student) error {
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
