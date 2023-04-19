package helpers

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// Checking if there is any error
func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Getting request body
func GetRequestBody(c *gin.Context) map[string]interface{} {
	var data map[string]interface{}
	var err = c.BindJSON(&data)
	CheckNilErr(err)
	return data
}

// Converting CSV to map[string]interface{}
func CsvToMap(csvFile *multipart.FileHeader) []interface{} {
	// Converting csv to 2D array
	file, err := csvFile.Open()
	CheckNilErr(err)
	var reader = csv.NewReader(file)
	array, err := reader.ReadAll()
	CheckNilErr(err)
	// Convering 2Darray to array of map
	var arr []interface{}
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
		var m = map[string]interface{}{}
		for i, v := range row {
			m[headers[i]] = v
		}
		arr = append(arr, m)
	}
	return arr
}
