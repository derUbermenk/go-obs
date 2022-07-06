package repository

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

	// assert that the file for the environment exist

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
