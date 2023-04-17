package main

import (
	"net/http"
	"os"

	"github.com/FulgurCode/school-erp-api/connections"
	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/FulgurCode/school-erp-api/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviornment variables
	godotenv.Load(".env")

	// Connecting to Database
	connections.ConnectDatabase()

	// Setting up gin router
	var router = router.Router()
	router.Run()

	// Listen to requests in 'port'
	var port = os.Getenv("PORT")
	var err = http.ListenAndServe(":"+port, router)
	helpers.CheckNilErr(err)

	// Disconnecting from Database
	defer connections.DisconnectDatabase()
}
