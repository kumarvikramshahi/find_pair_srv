package configs

type Configuration struct {
	Port    string
	EnvName string
}

// service config
var ServiceConfiguration Configuration

// LoadConfiguration for all env
func LoadConfiguration(fileName string) {

	switch configs := fileName; configs {

	case "dev":
		conf := Configuration{}
		devConfig := DevConfig{&conf}
		devConfig.ConfigManager()
		ServiceConfiguration = *devConfig.Configuration

	case "env":
		conf := Configuration{}
		prodConfig := EnvConfig{&conf}
		prodConfig.ConfigManager()
		ServiceConfiguration = *prodConfig.Configuration

	default:
		conf := Configuration{}
		localConfig := DevConfig{&conf}
		localConfig.ConfigManager()
		ServiceConfiguration = *localConfig.Configuration
	}
}
