package router

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// New gin Router
	var router = gin.Default()

	// Creating sessions
	var secret = os.Getenv("SECRET_KEY")
	var store = cookie.NewStore([]byte(secret))
	var sessionNames = []string{"admin"}
	router.Use(sessions.SessionsMany(sessionNames, store))

	// Admin routes
	AdminRouter(router.Group("/api/admin"))

	return router
}
