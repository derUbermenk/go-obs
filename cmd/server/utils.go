package main

import (
	"database/sql"
	"online-bidding-system/pkg/repository"
)

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
