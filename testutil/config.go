package testutil

import (
	"gitlab.com/trustify/core/config"
	"gitlab.com/trustify/core/pkg/util/environment"
)

// ReadConfig reads config file for test
func ReadConfig() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.Test,
	})
}
