package router

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// New gin Router
	var router = gin.New()

	// cors setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"https://www.chattanchalhss.com", "https://chattanchalhss.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Creating sessions
	var secret = os.Getenv("SECRET_KEY")
	var store = cookie.NewStore([]byte(secret))
	var sessionNames = []string{"admin", "teacherSignupOTP", "teacher"}
	router.Use(sessions.SessionsMany(sessionNames, store))

	// Admin routes
	AdminRouter(router.Group("/api/admin"))
	// Teacher routes
	TeacherRouter(router.Group("/api/teacher"))

	router.GET("/ws/admission-photo", func(c *gin.Context) {
		HandleWebSocket(c)
	})

	return router
}
