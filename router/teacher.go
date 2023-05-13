package router

import (
	"github.com/FulgurCode/school-erp-api/controller"
	"github.com/gin-gonic/gin"
)

func TeacherRouter(router *gin.RouterGroup) {
	// Authentication
	// signup route
	router.POST("/signup", controller.TeacherSignup)
	// signup-otp route
	router.GET("/signup-otp", controller.TeacherSignupOTP)
	// login route
	router.POST("/login", controller.TeacherLogin)
	// checklogin route
	router.GET("/checklogin", controller.TeacherCheckLogin)
}
