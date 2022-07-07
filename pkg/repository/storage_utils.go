package repository

import (
	"fmt"
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
	DatabaseName string "json:database_name"
}

func NewDatabaseConfig(env string) (databaseConfig *DatabaseConfig, err error) {

	// assert that the environment exists
	// err = AssertDatabaseConfigEnvExists(env)

	// get file path for the environment
	// file_path := GetConfigFilePath(env)

	// assert that the file for the environment exist
	// err := AssertConfigFileExists(file_path string)

	// unmarshall the contents of the file into database config

	// format the connstring for the database config

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
