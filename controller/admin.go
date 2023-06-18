package controller

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/helpers/adminHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/databaseHelpers"
	"github.com/FulgurCode/school-erp-api/helpers/studentHelpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// POST request on /api/admin/login
func AdminLoginRoute(c *gin.Context) {
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Checking for username
	var admin, err = databaseHelpers.GetAdminWithUsername(data["username"].(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(401, "Wrong username or password")
			return
		}
		c.JSON(500, "Request failed")
	}
	// Comparing password and sending response
	var result = helpers.ComparePassword(admin["password"].(string), data["password"].(string))
	if !result {
		c.JSON(401, "Wrong username or password")
		return
	}
	// storing id and sending response if password is correct
	adminHelpers.LoginWithSesssion(c, admin)
	c.JSON(200, "Login Successful")
}

// GET request on '/api/admin/checklogin'
func AdminCheckLoginRoute(c *gin.Context) {
	// checking if logged in as admin and sending response
	var isLoggedIn = adminHelpers.CheckLogin(c)
	c.JSON(200, isLoggedIn)
}

// DELETE request on '/api/admin/logout'
func AdminLogoutRoute(c *gin.Context) {
	// clearing 'admin' session
	adminHelpers.Logout(c)
	// Response for the request
	c.JSON(200, "Loggged Out")
}

// PUT request on '/api/admin/change-password'
func ChangeAdminPassword(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting request body
	var data = helpers.GetRequestBody(c)
	data["id"] = adminHelpers.GetId(c)
	// Checking if old password is correct
	var admin, _ = databaseHelpers.GetAdmin(data["id"].(primitive.ObjectID))
	var result = helpers.ComparePassword(admin["password"].(string), data["old-password"].(string))
	if !result {
		c.JSON(401, "Old password is wrong")
		return
	}
	// Changing admin password and sending response
	var err = adminHelpers.ChangePassword(data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Password Changed")
}

// POST request on '/api/admin/new-admission'
func NewAdmissionRoute(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting request body
	var data = helpers.GetStudentBody(c)
	// Adding student to database and sending response
	var id, err = studentHelpers.AddNewAdmission(data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, id)
}

// GET request on '/api/admin/get-students'
func GetStudentsRoute(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting search details
	var data = helpers.GetRequestBody(c)
	var name = c.Query("name")
	// Getting students and sending response
	var students, err = studentHelpers.GetStudents(data, name)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, students)
}

// POST request on '/api/admin/import-students'
func ImportStudents(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting uploaded data file
	var file, err = c.FormFile("file")
	helpers.CheckNilErr(err)
	var students = studentHelpers.ImportStudentsFromCSV(file)
	// importing students to database sending response
	err = databaseHelpers.ImportStudents(students)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Successfully added")
}

// PUT request on '/api/admin/edit-student'
func EditStudent(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting object id of student
	var studentId, err = primitive.ObjectIDFromHex(c.Query("studentId"))
	helpers.CheckNilErr(err)
	// Getting request body
	var data = helpers.GetRequestBody(c)
	err = studentHelpers.EditStudent(studentId, data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Updated Successfully")
}

// POST request on '/api/admin/upload-student-photo'
func UploadStudentPhoto(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting id of student
	var studentId = c.Query("studentId")
	// Get uploaded image
	var file, err = c.FormFile("file")
	helpers.CheckNilErr(err)
	err = c.SaveUploadedFile(file, "./public/images/students/"+studentId+".jpg")
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Successfully added")
}

// GET request on '/api/admin/get-student-photo'
func GetStudentPhoto(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting id of student
	var studentId = c.Query("studentId")
	// Getting image and sending response
	var file, err = ioutil.ReadFile("./public/images/students/" + studentId + ".jpg")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.JSON(404, "Image not found")
			return
		}
		c.JSON(500, "Request failed")
		return
	}
	var str = base64.StdEncoding.EncodeToString(file)
	c.JSON(200, str)
}

// POST request on '/api/admin/add-teacher'
func AddTeacher(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// Add teacher to database
	var err = databaseHelpers.AddTeacher(data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Teacher added")
}

// POST request on '/api/admin/add-duty'
func AddDuty(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting request body
	var data = helpers.GetRequestBody(c)
	// add new duty
	var err = databaseHelpers.AddDuty(data)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Duty added")
}

// GET request on '/api/admin/get-student'
func GetStudent(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting object id of student
	var studentId, err = primitive.ObjectIDFromHex(c.Query("studentId"))
	helpers.CheckNilErr(err)
	// Getting student using id
	student, err := databaseHelpers.GetStudent(studentId)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, student)
}

// GET request on '/api/admin/get-teacher'
func GetTeacher(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting object id of teacher
	var teacherId, err = primitive.ObjectIDFromHex(c.Query("teacherId"))
	helpers.CheckNilErr(err)
	// Getting teacher using id
	teacher, err := databaseHelpers.GetTeacher(teacherId)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, teacher)
}

// POST request on '/api/admin/import-teachers'
func ImportTeachers(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged In")
		return
	}
	// Getting uploaded data file
	var file, err = c.FormFile("file")
	helpers.CheckNilErr(err)
	var teachers = studentHelpers.ImportTeachersFromCSV(file)
	// importing teachers to database sending response
	err = databaseHelpers.ImportTeachers(teachers)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Successfully added")
}

// GET request on '/api/admin/get-teachers'
func GetTeachers(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// get teachers and send response
	var teachers, err = databaseHelpers.GetAllTeachers()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, teachers)
}

// GET request on '/api/admin/get-duties'
func GetDuties(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Get duties and send as response
	var duties, err = databaseHelpers.GetDuties()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, duties)
}

// DELETE requset on '/api/admin/delete-duty'
func DeleteDuty(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// deleting duty and sending response
	var duty, err = primitive.ObjectIDFromHex(c.Query("duty"))
	err = databaseHelpers.DeleteDuty(duty)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Duty deleted")
}

// GET request on '/api/admin/students-to-confirm'
func AdminStudentsToConfirm(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting students that is remain to confirm and sending response
	var students, err = databaseHelpers.GetStudentsToConfirm()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, students)
}

// PATCH request on '/api/admin/confirm-student'
func AdminConfirmStudent(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting student id
	var studentId, _ = primitive.ObjectIDFromHex(c.Query("studentId"))
	// verifying student and sending response
	var err = databaseHelpers.ConfirmStudent(studentId)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Student Confirmed")
}

// GET request on '/api/admin/students-to-verify'
func AdminStudentsToVerify(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting students that is remain to verify and sending response
	var students, err = databaseHelpers.GetStudentsToVerify()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, students)
}

// PATCH request on '/api/admin/verify-student'
func AdminVerifyStudent(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting student id
	var studentId, _ = primitive.ObjectIDFromHex(c.Query("studentId"))
	// verifying student and sending response
	var err = databaseHelpers.VerifyStudent(studentId)
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, "Student verifyed")
}

// Get request on '/api/admin/course-language-report'
func AdminCourseLanguageReport(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting data from database and sending response
	var data, err = databaseHelpers.CourseLanguageReport()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, data)
}

// Get request on '/api/admin/course-status-report'
func AdminCourseStatusReport(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting data from database and sending response
	var data, err = databaseHelpers.CourseStatusReport()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, data)
}

// Get request on '/api/admin/course-gender-report'
func AdminCourseGenderReport(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting data from database and sending response
	var data, err = databaseHelpers.CourseGenderReport()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, data)
}

// Get request on '/api/admin/course-category-report'
func AdminCourseCategoryReport(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting data from database and sending response
	var data, err = databaseHelpers.CourseCategoryReport()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, data)
}

// Get request on '/api/admin/course-caste-report'
func AdminCourseCasteReport(c *gin.Context) {
	// Checking if logged in
	if !adminHelpers.CheckLogin(c) {
		c.JSON(401, "Not Logged in admin")
		return
	}
	// Getting data from database and sending response
	var data, err = databaseHelpers.CourseCasteReport()
	if err != nil {
		c.JSON(500, "Request failed")
		return
	}
	c.JSON(200, data)
}
