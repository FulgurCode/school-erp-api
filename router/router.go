package router

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Cors widdleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func Router() *gin.Engine {
	// New gin Router
	var router = gin.Default()
	router.Use(CORSMiddleware())

	// Creating sessions
	var secret = os.Getenv("SECRET_KEY")
	var store = cookie.NewStore([]byte(secret))
	var sessionNames = []string{"admin"}
	router.Use(sessions.SessionsMany(sessionNames, store))

	// Admin routes
	AdminRouter(router.Group("/api/admin"))

	return router
}
