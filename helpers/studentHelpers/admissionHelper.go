package studentHelpers 

import (
	"sync"

	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WaitGroup
var wg sync.WaitGroup

// Get admission number and add new admission to database
func AddNewAdmission(student map[string]interface{}) (string, error) {
	// waiting for request to finish
	wg.Wait()
	wg.Add(1)
	// getting last admission number
	if _, exists := student["admissionNo"]; !exists && student["status"] == "permanent" {
		var admissionNo = databaseHelpers.GetLastAdmissionNumber()
		student["admissionNo"] = admissionNo + 1
	}
	// inserting student to database
	var id, err = databaseHelpers.InsertStudent(student)
	wg.Done()
	return id, err
}

// Set admission number if student don't have one and update student details
func EditStudent(studentId primitive.ObjectID, student map[string]interface{}) error {
	// waiting for request to finish
	wg.Wait()
	wg.Add(1)
	// getting last admission number
	if _, exists := student["admissionNo"]; !exists && student["status"] == "permanent" {
		var admissionNo = databaseHelpers.GetLastAdmissionNumber()
		student["admissionNo"] = admissionNo + 1
	}
	// Update student details
	var err = databaseHelpers.UpdateStudent(studentId, student)
	wg.Done()
	return err
}
