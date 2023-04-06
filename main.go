package main

import (
	"net/http"
	"os"

	"github.com/FulgurCode/school-erp-api/connections"
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/router"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load enviornment variables
	godotenv.Load(".env")

	// Connecting to Database
	connections.ConnectDatabase()

	// Setting up gin router
	var router = router.Router()
	router.Run()

	// cors setup
	var c = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	})

	// Listen to requests in 'port'
	var port = os.Getenv("PORT")
	var err = http.ListenAndServe(":"+port, c.Handler(router))
	helpers.CheckNilErr(err)

	// Disconnecting from Database
	defer connections.DisconnectDatabase()
}
