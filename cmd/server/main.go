package main

import (
	"fmt"
	"log"
	"online-bidding-system/cmd/serverutils"
	"online-bidding-system/pkg/app"
	"online-bidding-system/pkg/repository"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %s \\n", err)
		os.Exit(1)
	}
}

// Runs the app. The follwing steps are taken.
//   1. initialize a DatabaseConfig
//   2. format the connection
//   3. Connect to the database using the connection string
//   4. Initialize a Storage interface type for use of apis
//   5. initialize a router for use of the server
//   6. Setup the server with the dependencies, router and apis as arguments
func run() error {

	// initialize DatabaseConfig
	// DatabaseConfig reads json config files
	// that contain information needed for connecting to database
	// See DatabaseConfig for more info.
	dev_env := "dev"
	dbconfig, err := repository.NewDatabaseConfig(dev_env)

	if err != nil {
		return err
	}

	// The connectionstring is required input when working with
	// database level code.
	connection_string := dbconfig.ConnectionString()

	// Connect to the database specified by the connection string
	db, err := serverutils.ConnectDatabase(connection_string)

	if err != nil {
		log.Printf("Err on Main SetupStorage: %v\n", err)
		return err
	}

	// initializes a Storage type variable for use as dependencies of apis
	// and subsequently run migrations on the storage's database.
	//
	// * replace _ to storage when ready
	storage, err = serverutils.SetupStorage(connection_string, db)

	if err != nil {
		log.Printf("Err on Main SetupStorage: %v\n", err)
		return err
	}

	// setup the api to be used by the server
	//
	// typical api setup code.
	// api_1 := api.NewApi1(storage)
	user_service := api.NewUserService(storage)
	bidding_service := api.NewBiddingService(storage)
	auth_service := api.NewAuthService(storage)

	// server setup
	// we add our routes to the router
	router := gin.Default()
	router.Use(cors.Default())

	// the router is a dependency of the server
	server := app.NewServer(router, user_service, bidding_service, auth_service)

	// run the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}
