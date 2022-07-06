package main

import (
	"database/sql"
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

	var connection_string string

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

func loadDatabaseConfig(config *DatabaseConfig) {
	file, err := os.Open("../../pkg/repository/configs/development_db")
}

func setupStorage(connection_string string, db *sql.DB) (storage repository.Storage, err error) {
	storage = repository.NewStorage(db)
	err = storage.RunMigrations(connection_string)

	if err != nil {
		return nil, err
	}

	return storage, nil
}

func connectDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
