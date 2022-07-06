package repository_test

import (
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
			want_err: repository.ErrNonExistentDatabaseEnvironment,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := repository.AssertDatabaseConfigEnvExists(test.env)

			if err != nil {
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
