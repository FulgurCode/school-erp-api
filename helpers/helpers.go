package helpers

import "github.com/gin-gonic/gin"

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
