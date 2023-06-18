package component_tests

import (
	"fmt"
	"github.com/brimb0r-org/configurator/configurator"
	"github.com/brimb0r-org/eq/application/internal/internal_config"
	"os"
)

var (
	ComponetTests = os.Getenv("COMPONET_TESTS")
)

func configureTests() *internal_config.Configuration {
	configPath := fmt.Sprintf("../config_files/")
	_ = os.Setenv("CONFIG_PATH", configPath)
	configure := configurator.New()
	configuration := &internal_config.Configuration{}
	configure.Unmarshal(configuration)
	return configuration
}
