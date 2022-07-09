package repository

import (
	"fmt"
	"reflect"
)

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

type ErrNonExistentConfigFile struct {
	Filepath string
}

func (e *ErrNonExistentConfigFile) Error() string {
	msg := fmt.Sprintf("Non existent config file: \n\t%v", e.Filepath)
	return msg
}

func (e *ErrNonExistentConfigFile) Is(target error) bool {
	t, isType := target.(*ErrNonExistentConfigFile)

	if !isType {
		return false
	}

	return t.Filepath == e.Filepath
}

type ErrParsingFile struct {
	Err error
}

func (e *ErrParsingFile) Error() string {
	msg := fmt.Sprintf("Error parsing file: %v", e.Err)
	return msg
}

func (e *ErrParsingFile) Is(target error) bool {
	t, isType := target.(*ErrParsingFile)

	return isType && t.Err == e.Err
}

type ErrIncompleteDatabaseConfig struct {
	MissingUsername     bool
	MissingPassword     bool
	MissingHost         bool
	MissingDatabaseName bool
}

func (e *ErrIncompleteDatabaseConfig) Error() string {
	// iterate over the error fields

	base_msg := "The following fields are not present in the configuration: "

	// for every error field that is true
	// 	base_msg := fmt.Sprintf("%v%v", base_msg, missing_field_format)
	// add it to the message

	v := reflect.ValueOf(e)
	typeofE := v.Type()

	for i := 0; i < v.NumField(); i++ {
		missing_field := typeofE.Field(i).Name
		is_missing_field := v.Field(i).Interface().(bool)

		if is_missing_field {
			missing_field_format := fmt.Sprintf("\n\t%v", missing_field)
			base_msg = fmt.Sprintf(base_msg, missing_field_format)
		}
	}

	return base_msg
}

func (e *ErrIncompleteDatabaseConfig) Is(target error) bool {
	t, isType := target.(*ErrIncompleteDatabaseConfig)

	if !isType {
		return false
	}

	return (e.MissingUsername == t.MissingUsername &&
		e.MissingPassword == t.MissingPassword &&
		e.MissingHost == t.MissingHost &&
		e.MissingDatabaseName == t.MissingDatabaseName)
}
