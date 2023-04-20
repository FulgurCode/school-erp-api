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
	router.GET("/logout", controller.AdminLogoutRoute)
	// change-password route
	router.PUT("/change-password", controller.ChangeAdminPassword)

	// Admission
	// new-admission route
	router.POST("/new-admission", controller.NewAdmissionRoute)
  // import-students route
  router.POST("/import-students", controller.ImportStudents)

	// Student
	// get-student route
	router.GET("/get-students", controller.GetStudentsRoute)
}
