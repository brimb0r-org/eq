package internal_config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_LocalConfigIntegration(t *testing.T) {
	if os.Getenv("COMPONENT") == "" {
		t.Skip()
	}
	env := os.Getenv("ENVIRONMENT")
	configPath := os.Getenv("CONFIG_PATH")

	defer func() {
		os.Setenv("ENVIRONMENT", env)
		os.Setenv("CONFIG_PATH", configPath)
	}()

	os.Setenv("ENVIRONMENT", "local")
	os.Setenv("CONFIG_PATH", "../../config_files/")
	configuration := Configure()

	assert.Equal(t, "local", configuration.Environment)
	assert.Equal(t, "eq", configuration.ServiceName)
}
