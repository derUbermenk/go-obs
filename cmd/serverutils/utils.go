package serverutils

import (
	"database/sql"
	"online-bidding-system/pkg/repository"

	_ "github.com/lib/pq"
)

// initializes a storage type and runs migrations on its database field.
func SetupStorage(connection_string string, db *sql.DB) (storage repository.Storage, err error) {
	storage = repository.NewStorage(db)
	err = storage.RunMigrations(connection_string)

	if err != nil {
		return nil, err
	}

	return storage, nil
}

// connects to a database. This allows any storage type with the field type database
// do actions on it.
func ConnectDatabase(connString string) (*sql.DB, error) {
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
