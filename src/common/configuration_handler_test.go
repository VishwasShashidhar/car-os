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

		for key, value := range keyValueSet {
			os.Unsetenv(key)
			actual := GetConfigurationForKey(key)
			assert.Equal(t, value, actual)
		}
	})

}
