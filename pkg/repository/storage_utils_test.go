package repository_test

import (
	"errors"
	"online-bidding-system/pkg/repository"
	"testing"
)

func TestConnectionString(t *testing.T) {
	tests := []struct {
		name                  string
		databaseConfig        *repository.DatabaseConfig
		want_connectionString string
	}{
		{
			name: "Returns the correct connection string for the following database config",
			databaseConfig: &repository.DatabaseConfig{
				Username:     "test_user",
				Password:     "test_password",
				Host:         "test_host",
				DatabaseName: "test_database",
			},
			want_connectionString: "postgres://test_user:test_password@test_host/test_database?sslmode=disable",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			connString := test.databaseConfig.ConnectionString()

			if connString != test.want_connectionString {
				t.Errorf(
					"\nTest failed: %v\n\tgot: %v\n\twant: %v",
					test.name,
					connString,
					test.want_connectionString,
				)
			}
		})
	}
}

func TestAssertDatabaseConfigEnvExists(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		want_err error
	}{
		{
			name:     "Returns no error when given existent environment v1",
			env:      "dev",
			want_err: nil,
		},
		{
			name:     "Returns no error when given existent environment v2",
			env:      "test",
			want_err: nil,
		},
		{
			name:     "Returns no error when given non existent environment",
			env:      "missingEnv",
			want_err: &repository.ErrNonExistentDatabaseEnvironment{Env: "missingEnv"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := repository.AssertDatabaseConfigEnvExists(test.env)

			if !errors.Is(err, test.want_err) {
				t.Errorf(
					"Test failed: %v \n\tgot: %v\n\twant: %v",
					test.name,
					err,
					test.want_err,
				)
			}
		})
	}
}

func TestGetConfigFilePath(t *testing.T) {
	tests := []struct {
		name          string
		env           string
		want_filepath string
	}{
		{
			name:          "When given the env it returns the correct file path v1",
			env:           "test",
			want_filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_test.json",
		},
		{
			name:          "When given the env it returns the correct file path v2",
			env:           "dev",
			want_filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_dev.json",
		},
		{
			name:          "When given the env it returns the correct file path v3",
			env:           "some_env",
			want_filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_some_env.json",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file_path := repository.GetConfigFilePath(test.env)

			if file_path != test.want_filepath {
				t.Errorf(
					"Test failed: %v \n\tgot %v \n\twant: %v",
					test.name,
					file_path,
					test.want_filepath,
				)
			}
		})
	}
}

func TestAssertConfigFileExists(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want_err error
	}{
		{
			name:     "Returns no error when given an existing database config file v1",
			filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_test.json",
			want_err: nil,
		},
		{
			name:     "Returns no error when given an existing database config file v2",
			filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_dev.json",
			want_err: nil,
		},
		{
			name:     "Returns an error when given an non existing filepath",
			filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_nonExisting.json",
			want_err: &repository.ErrNonExistentConfigFile{Filepath: "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/obs_nonExisting.json"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := repository.AssertConfigFileExists(test.filepath)

			if !errors.Is(err, test.want_err) {
				t.Errorf(
					"Test failed: %v \n\tgot: %v\n\twant: %v",
					test.name,
					err,
					test.want_err,
				)
			}
		})
	}
}

func TestParseConfigFile(t *testing.T) {
	tests := []struct {
		name                string
		databaseConfig      *repository.DatabaseConfig
		file_path           string
		want_databaseConfig *repository.DatabaseConfig
		want_err            error
	}{
		{
			name:           "Parses the database config",
			databaseConfig: &repository.DatabaseConfig{},
			file_path:      "/home/chester/Documents/code_projects/go-projects/go-online-bidding-system/pkg/repository/configurations/sample_db.json",
			want_databaseConfig: &repository.DatabaseConfig{
				Username:     "test_user",
				Password:     "test_password",
				Host:         "test_local",
				DatabaseName: "test_db",
			},
			want_err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := repository.ParseConfigFile(test.databaseConfig, test.file_path)

			if !errors.Is(err, test.want_err) {
				t.Errorf(
					"Test failed: %v \n\tgot: %v\n\twant: %v",
					test.name,
					err,
					test.want_err,
				)
			}

			if *test.databaseConfig != *test.want_databaseConfig {
				t.Errorf(
					"Test failed: %v \n\tgot: %v\n\twant: %v",
					test.name,
					test.databaseConfig,
					test.want_databaseConfig,
				)
			}
		})
	}
}

func AssertCompleteConfig(t *testing.t) {
	tests := []struct {
		name           string
		databaseConfig *repository.DatabaseConfig
		want_err       error
	}{
		{
			name:           "Returns an error when given an incomplete database config. v1",
			databaseConfig: &repository.DatabaseConfig{},
			want_err: &repository.ErrIncompleteDatabaseConfig{
				MissingUsername:     true,
				MissingPassword:     true,
				MissingHost:         true,
				MissingDatabaseName: true,
			},
		},
		{
			name: "Returns an error when given an incomplete database config. v2",
			databaseConfig: &repository.DatabaseConfig{
				Password:     "pw",
				Host:         "s",
				DatabaseName: "s",
			},
			want_err: &repository.ErrIncompleteDatabaseConfig{
				MissingUsername: true,
			},
		},
		{
			name: "Returns an error when given an incomplete database config. v2",
			databaseConfig: &repository.DatabaseConfig{
				Username:     "s",
				Host:         "s",
				DatabaseName: "s",
			},
			want_err: &repository.ErrIncompleteDatabaseConfig{
				MissingPassword: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t testing.T) {
			err := repository.AssertCompleteConfig(test.databaseConfig)

			if !errors.Is(test.want_err, err) {
				t.Errorf(
					"\nTest failed: %v\n\tgot: %v\n\twant: %v",
					test.name,
					err,
					test.want_err,
				)
			}
		})
	}
}

func TestNewDataBaseConfig(t *testing.T) {
	no_err_tests := []struct {
		name     string
		env      string
		want_err error
	}{
		{
			name:     "Returns a nil error when given a valid environment. dev",
			env:      "dev",
			want_err: nil,
		},
		{
			name:     "Returns a nil error when given a valid environment. test",
			env:      "test",
			want_err: nil,
		},
	}

	for _, test := range no_err_tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := repository.NewDatabaseConfig(test.env)

			if !errors.Is(test.want_err, err) {
				t.Errorf(
					"\nTest failed: %v \n\tgot: %v\n\twant: %v",
					test.name,
					err,
					test.want_err,
				)
			}
		})
	}

	filled_fields_tests := []struct {
		name string
		env  string
	}{
		{
			name: "Returns a valid DataConfig with all fields filled with env dev",
			env:  "dev",
		},
		{
			name: "Returns a valid DataConfig with all fields filled with env test",
			env:  "test",
		},
	}

	for _, test := range filled_fields_tests {
		t.Run(test.name, func(t *testing.T) {
			databaseConfig, _ := repository.NewDatabaseConfig(test.env)

			if databaseConfig.Username == "" {
				t.Errorf(
					"\nTest failed: %v\n\t databaseConfig Username empty",
					test.name,
				)
			}

			if databaseConfig.Password == "" {
				t.Errorf(
					"\nTest failed: %v\n\t databaseConfig Password empty",
					test.name,
				)
			}

			if databaseConfig.Host == "" {
				t.Errorf(
					"\nTest failed: %v\n\t databaseConfig Host empty",
					test.name,
				)
			}

			if databaseConfig.DatabaseName == "" {
				t.Errorf(
					"\nTest failed: %v\n\t databaseConfig DatabaseName empty",
					test.name,
				)
			}
		})
	}

	/*
		{
			name: "Returns a valid DataConfig with the valid database for the environment v1",
			env: "dev",
			want_databaseConfigDatabaseName: "obs",
		},
		{
			name: "Returns a valid DataConfig with the valid database for the environment v2",
			env: "test",
			want_databaseConfigDatabaseName: "obs_test",
		},
	*/

	valid_database_name_tests := []struct {
		name              string
		env               string
		want_databasename string
	}{
		{
			name:              "Returns the correct database name for the environment. dev",
			env:               "dev",
			want_databasename: "obs",
		},
		{
			name:              "Returns the correct database name for the environment. test",
			env:               "test",
			want_databasename: "obs_test",
		},
	}

	for _, test := range valid_database_name_tests {
		t.Run(test.name, func(t *testing.T) {
			databaseConfig, _ := repository.NewDatabaseConfig(test.env)
			databasename := databaseConfig.DatabaseName

			if test.want_databasename != databasename {
				t.Errorf(
					"Test failed: %v \n\tgot: %v\n\twant: %v",
					test.name,
					databasename,
					test.want_databasename,
				)
			}
		})
	}
}
