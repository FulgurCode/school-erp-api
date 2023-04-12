package adminHelpers

import (
	"sync"

	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
)

// WaitGroup
var wg sync.WaitGroup

// Get admission number and add new admission to database
func AddNewAdmission(student map[string]interface{}) error {
	// waiting for request to finish
	wg.Wait()
	wg.Add(1)
	// getting last admission number
	var admissionNo = databaseHelpers.GetLastAdmissionNumber()
	student["admissionNo"] = admissionNo + 1
	// inserting student to database
	var err = databaseHelpers.InsertStudent(student)
	wg.Done()
	return err
}
