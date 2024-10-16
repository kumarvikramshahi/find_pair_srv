package configs

// DevConfig struct
type DevConfig struct {
	*Configuration
}

// ConfigManager to mange the config details for Locals
func (conf *DevConfig) ConfigManager() *Configuration {
	conf.Port = "8000"
	conf.EnvName = "local"

	return conf.Configuration
}
