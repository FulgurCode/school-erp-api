package router

import (
	"github.com/FulgurCode/school-erp-api/controller"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.RouterGroup) {
	// Authentication
	// login route
	router.POST("/login", controller.AdminLoginRoute)
	// checklogin route
	router.GET("/checklogin", controller.AdminCheckLoginRoute)
	// logout route
	router.DELETE("/logout", controller.AdminLogoutRoute)
	// change-password route
	router.PUT("/change-password", controller.ChangeAdminPassword)

	// Admission
	// new-admission route
	router.POST("/new-admission", controller.NewAdmissionRoute)
	// import-students route
	router.POST("/import-students", controller.ImportStudents)
	// edit-student route
	router.PUT("/edit-student", controller.EditStudent)
	// upload-student-photo route
	router.POST("/upload-student-photo", controller.UploadStudentPhoto)

	// Student
	// get-student route
	router.GET("/get-students", controller.GetStudentsRoute)

	// Teacher
	// add-teacher route
	router.POST("/add-teacher", controller.AddTeacher)
}
