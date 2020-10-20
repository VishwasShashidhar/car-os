package common

import (
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

func TestGetConfigurationForKey(t *testing.T) {

	Init()

	t.Run("test configuration values", func(t *testing.T) {
		keyValueSet := make(map[string]string)

		keyValueSet["APP_NAME"] = "Car OS Test"
		keyValueSet["APP_PORT"] = "8090"
		keyValueSet["DB_HOST"] = "127.0.0.1"
		keyValueSet["DB_PORT"] = "5432"
		keyValueSet["DB_USER"] = "postgres"
		keyValueSet["DB_PASSWORD"] = "test"
		keyValueSet["DB_DATABASE"] = "carostest"

		for key, value := range keyValueSet {
			os.Setenv(key, value)
			actual := GetConfigurationForKey(key)
			assert.Equal(t, value, actual)
		}
	})

	t.Run("test default values", func(t *testing.T) {
		keyValueSet := make(map[string]string)

		keyValueSet["APP_NAME"] = "CarOS"
		keyValueSet["APP_PORT"] = "8080"
		keyValueSet["DB_HOST"] = "127.0.0.1"
		keyValueSet["DB_PORT"] = "5432"
		keyValueSet["DB_USER"] = "john"
		keyValueSet["DB_PASSWORD"] = "doe"
		keyValueSet["DB_DATABASE"] = "caros"

		for key, value := range keyValueSet {
			os.Unsetenv(key)
			actual := GetConfigurationForKey(key)
			assert.Equal(t, value, actual)
		}
	})

}
