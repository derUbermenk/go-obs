package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var database_environment = map[string]string{
	"dev":  "dev",
	"test": "test",
}

type DatabaseConfig struct {
	Username     string "json:username"
	Password     string "json:password"
	Host         string "json:host"
	DatabaseName string "json:databasename"
}

func (dc *DatabaseConfig) ConnectionString() string {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		dc.Username,
		dc.Password,
		dc.Host,
		dc.DatabaseName,
	)

	return connString
}

func NewDatabaseConfig(env string) (databaseConfig *DatabaseConfig, err error) {
	databaseConfig = &DatabaseConfig{}

	// assert that the environment exists
	err = AssertDatabaseConfigEnvExists(env)

	if err != nil {
		return nil, err
	}

	// get file path for the environment
	file_path := GetConfigFilePath(env)

	// assert that the file for the environment exist
	err = AssertConfigFileExists(file_path)
	if err != nil {
		return nil, err
	}

	// unmarshall the contents of the file into database config
	err = ParseConfigFile(databaseConfig, file_path)

	if err != nil {
		return nil, err
	}

	// return database config
	return databaseConfig, nil
}

func AssertDatabaseConfigEnvExists(env string) (err error) {
	_, env_exists := database_environment[env]

	if !env_exists {
		err = &ErrNonExistentDatabaseEnvironment{
			Env: env,
		}
		return err
	}

	return nil
}

func GetConfigFilePath(env string) (file_path string) {
	_, callingFile, _, _ := runtime.Caller(0)
	file_path = filepath.Join(filepath.Dir(callingFile), "configurations", fmt.Sprintf("obs_%v.json", env))
	return
}

func AssertConfigFileExists(file_path string) error {
	file, err := os.Open(file_path)
	defer file.Close()

	if err != nil {
		return &ErrNonExistentConfigFile{
			Filepath: file_path,
		}
	}

	return nil
}

func ParseConfigFile(databaseConfig *DatabaseConfig, file_path string) error {
	file, err := os.Open(file_path)
	defer file.Close()

	if err != nil {
		err = &ErrParsingFile{
			Err: err,
		}

		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		err = &ErrParsingFile{
			Err: err,
		}
	}

	json.Unmarshal(data, databaseConfig)
	return nil
}
