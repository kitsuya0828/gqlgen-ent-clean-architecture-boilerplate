package testutil

import "github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/config"

// ReadConfig reads config file for test.
func ReadConfig() {
	config.ReadConfig(config.ReadConfigOption{
	   AppEnv: config.Test,
	})
 }