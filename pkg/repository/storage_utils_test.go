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
			want_err: &repository.ErrNonExistentConfigFile{Filepath: ""},
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
