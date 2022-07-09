package main

import (
	"fmt"
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

func run() error {
	// get database connection string
	dev_env := "dev"
	dbconfig, err := repository.NewDatabaseConfig(dev_env)

	if err != nil {
		return err
	}

	connection_string := dbconfig.ConnectionString()

	// database setup
	db, err := connectDatabase(connection_string)

	if err != nil {
		return err
	}

	_, err = setupStorage(connection_string, db)

	if err != nil {
		return err
	}

	// api setup

	// server setup
	router := gin.Default()
	router.Use(cors.Default())

	server := app.NewServer(router)

	// run the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}
