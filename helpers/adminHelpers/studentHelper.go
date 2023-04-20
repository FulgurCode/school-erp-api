package adminHelpers

import (
	"strconv"

	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
)

// Getting students according to search type
func GetStudents(search string, value string) ([]map[string]interface{}, error) {
	switch search {
	case "admissionNo":
		// Getting student by admission number
		var admissionNo, _ = strconv.Atoi(value)
		var students, err = databaseHelpers.GetStudentByAdmissionNo(admissionNo)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "applicationNo":
		// Getting student by application number
		var applicationNo, _ = strconv.Atoi(value)
		var students, err = databaseHelpers.GetStudentByApplicationNo(applicationNo)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "name":
		// Getting student by name
		var name = value
		var students, err = databaseHelpers.GetStudentByName(name)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	default:
		// Sending empty array if search type is random
		return []map[string]interface{}{}, nil
	}
}
