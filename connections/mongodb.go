package connections

import (
	"context"
	"os"

	"github.com/FulgurCode/school-erp-api/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Db *mongo.Database

// Connecting to mongodb Database
func ConnectDatabase() {
	// getting mongodb uri
	var uri = os.Getenv("MONGODB_URI")
	// Connection to database
	var ctx = context.Background()
	var options = options.Client().ApplyURI(uri)
	var client, err = mongo.Connect(ctx, options)
	helpers.CheckNilErr(err)

	// Database
	Db = client.Database("school-erp")
}

// Disconnecting from Database
func DisconnectDatabase() {
	var ctx = context.Background()
	Client.Disconnect(ctx)
}
