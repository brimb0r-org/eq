package internal_config

import (
	"fmt"
	"gSheets/application/pkg/eq_aws"
	"gSheets/application/pkg/eq_aws/sstore"
	mongo "gSheets/application/pkg/eq_mongo"
	"gSheets/application/pkg/excel"
	"github.com/brimb0r-org/configurator/configurator"
)

type Configuration struct {
	Environment string `yaml:"environment"`
	Region      string `yaml:"awsRegion"`
	ServiceName string `yaml:"serviceName"`
	AwsAccount  string `yaml:"awsAccount"`
	Schedule    int    `yaml:"scheduleIntervalSeconds"`
	Excel       excel.ExcelConfig
	Mongo       mongo.Config
}

func Configure() *Configuration {
	configure := configurator.New()
	configure.SetAccessor("SECRET", sstore.NewSStoreClient(eq_aws.Session()))
	configuration := &Configuration{}
	err := configure.Unmarshal(configuration)
	if err != nil {
		panic(fmt.Errorf("unmarshalling error [%w]", err))
	}
	return configuration
}
