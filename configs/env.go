package configs

import "os"

// EnvConfig struct
type EnvConfig struct {
	*Configuration
}

// ConfigManager for prod
func (conf *EnvConfig) ConfigManager() *Configuration {
	conf.Port = os.Getenv("PORT")
	conf.EnvName = os.Getenv("ENV_NAME")

	return conf.Configuration
}
