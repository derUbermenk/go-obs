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

// initializes a DatabaseConfig struct from a json file
// assertions of field completeness are made before final initializitations
// and will return an error if an error was met instead.
//
// The following assertions are made before initializing
// the config file
//   1. wanted environment exists
//   2. config file for the environment exists
// 	 3. all the needed fields for the config exists
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

	// assert that all database config fields are filled
	err = AssertCompleteConfig(databaseConfig)

	if err != nil {
		return nil, err
	}

	// return database config
	return databaseConfig, nil
}

// checks if the given environment is present in the choice
// of database environments
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

// Gets the config file path given the environment of choice
// calling file returns the filename that is calling this particular function.
// GetConfigFilePath via runtime.Caller(0)
// filepath.Dir(callingFile) returns the directory of the calling file
//
// with filepath.Join(directory_of_calling_file, configuration_folder_name, file_name)
// we get
// file_path = "directory_of_calling_file/configurations/obs_environment.json"
// wherein the directory of calling file starts at /home assuming a unix ecosystem
func GetConfigFilePath(env string) (file_path string) {
	_, callingFile, _, _ := runtime.Caller(0)
	file_path = filepath.Join(filepath.Dir(callingFile), "configurations", fmt.Sprintf("obs_%v.json", env))
	return
}

// Checks if the config file exists given the file path
// this assertion is called for better errors
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

// Parses the config file fields to a DatabaseConfig object
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

// Asserts that the returned DatabaseConfig has its fields all filled.
// Assertion is important so that problems are met at DatabaseConfig
// initialization.
func AssertCompleteConfig(databaseConfig *DatabaseConfig) error {
	err := &ErrIncompleteDatabaseConfig{}

	// check first if there are no negative values for all the databasecOnfig fields
	noUsername := databaseConfig.Username == ""
	noPassword := databaseConfig.Password == ""
	noHost := databaseConfig.Host == ""
	noDatabaseName := databaseConfig.DatabaseName == ""

	if noUsername || noPassword || noHost || noDatabaseName {
		if noUsername {
			err.MissingUsername = noUsername
		}

		if noPassword {
			err.MissingPassword = noPassword
		}

		if noHost {
			err.MissingHost = noHost
		}

		if noDatabaseName {
			err.MissingDatabaseName = noDatabaseName
		}

		return err
	}

	return nil
}
