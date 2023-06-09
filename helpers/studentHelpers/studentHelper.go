package studentHelpers

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
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
		// Adding custom detail to student
		student["status"] = "pending"
		student["import"] = true
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
func GetStudents(data map[string]interface{}, name string) ([]map[string]interface{}, error) {
	// Getting student details with name
	if name == "" {
		var students, err = databaseHelpers.GetStudents(data)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	}
	// Getting student details without name
	var students, err = databaseHelpers.GetStudentsByName(name, data)
	if students == nil {
		return []map[string]interface{}{}, nil
	}
	return students, err
}
