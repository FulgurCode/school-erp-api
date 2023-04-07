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
}
