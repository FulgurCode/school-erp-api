package studentHelpers

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"go.mongodb.org/mongo-driver/bson"
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
			if number, err := strconv.Atoi(v); err != nil {
				student[headers[i]] = v
			} else {
				student[headers[i]] = number
			}
		}
		// Adding pending status to student
		student["status"] = "pending"
		// Appending student to students array
		students = append(students, student)
	}
	return students
}

// Converting CSV to map[string]interface{}
func ImportTeachersFromCSV(csvFile *multipart.FileHeader) []interface{} {
	// Converting csv to 2D array
	file, err := csvFile.Open()
	helpers.CheckNilErr(err)
	var reader = csv.NewReader(file)
	array, err := reader.ReadAll()
	helpers.CheckNilErr(err)
	// Convering 2Darray to array of map
	var teachers []interface{}
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
		var teacher = map[string]interface{}{}
		for i, v := range row {
			if number, err := strconv.Atoi(v); err != nil {
				teacher[headers[i]] = v
			} else {
				teacher[headers[i]] = number
			}
		}
		// Adding pending status to teacher
		teacher["status"] = "pending"
		// Appending teacher to teachers array
		teachers = append(teachers, teacher)
	}
	return teachers
}
// Getting students according to search type
func GetStudents(search string, value string, status string) ([]map[string]interface{}, error) {
	switch search {
	case "admissionNo":
		// Getting student by admission number
		var admissionNo, _ = strconv.Atoi(value)
		var query = bson.M{"admissionNo": admissionNo}
		var students, err = databaseHelpers.GetStudents(query)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "applicationNo":
		// Getting student by application number
		var applicationNo, _ = strconv.Atoi(value)
		var query bson.M
		if status != "" {
			query = bson.M{"applicationNo": applicationNo, "status": status}
		} else {
			query = bson.M{"applicationNo": applicationNo}
		}
		var students, err = databaseHelpers.GetStudents(query)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "name":
		// Getting student by name
		var name = value
		var students, err = databaseHelpers.GetStudentsByName(name, status)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	default:
		// Sending empty array if search type is random
		return []map[string]interface{}{}, nil
	}
}
