package studentHelpers

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/FulgurCode/school-erp-api/helpers"
)

// Converting CSV to map[string]interface{}
func ImportStudentsFromCSV(csvFile *multipart.FileHeader) []interface{} {
	// Converting csv to 2D array
	file, err := csvFile.Open()
	helpers.CheckNilErr(err)
	var reader = csv.NewReader(file)
	array, err := reader.ReadAll()
	helpers.CheckNilErr(err)
	// Convering 2Darray to array of map
	var students []interface{}
	var headers []string
	for index, row := range array {
		// Getting key for map
		if index == 0 {
			for i := 0; i < len(row); i++ {
				headers = append(headers, row[i])
			}
			continue
		}
		// Getting valuse for the map
		var student = map[string]interface{}{}
		for i, v := range row {
			student[headers[i]] = v
		}
		// Adding pending status to student
		student["status"] = "pending"
		// Appending student to students array
		students = append(students, student)
	}
	return students
}
