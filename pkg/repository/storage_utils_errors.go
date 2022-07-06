package repository

import "fmt"

type ErrNonExistentDatabaseEnvironment struct {
	Env string
}

func (e *ErrNonExistentDatabaseEnvironment) Error() string {
	msg := fmt.Sprintf("Non existent database env: %v", e.Env)
	return msg
}

func (e *ErrNonExistentDatabaseEnvironment) Is(target error) bool {
	t, isType := target.(*ErrNonExistentDatabaseEnvironment)

	if !isType {
		return false
	}

	return t.Env == e.Env
}
