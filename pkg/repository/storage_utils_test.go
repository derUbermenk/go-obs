package repository_test

import (
	"errors"
	"online-bidding-system/pkg/repository"
	"testing"
)

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
			file_path:      "",
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

			if test.databaseConfig != test.want_databaseConfig {
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
