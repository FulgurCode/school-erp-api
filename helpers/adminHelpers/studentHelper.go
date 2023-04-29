package adminHelpers

import (
	"strconv"

	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"go.mongodb.org/mongo-driver/bson"
)

// Getting students according to search type
func GetStudents(search string, value string, status string) ([]map[string]interface{}, error) {
	switch search {
	case "admissionNo":
		// Getting student by admission number
		var admissionNo, _ = strconv.Atoi(value)
		var query = bson.M{"admissionNo": admissionNo}
		var students, err = databaseHelpers.GetStudentByAdmissionNo(query)
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
		var students, err = databaseHelpers.GetStudentByApplicationNo(query)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	case "name":
		// Getting student by name
		var name = value
		var query bson.M
		if status != "" {
			query = bson.M{"name": name, "status": status}
		} else {
			query = bson.M{"name": name}
		}
		var students, err = databaseHelpers.GetStudentByName(query)
		if students == nil {
			return []map[string]interface{}{}, nil
		}
		return students, err
	default:
		// Sending empty array if search type is random
		return []map[string]interface{}{}, nil
	}
}
