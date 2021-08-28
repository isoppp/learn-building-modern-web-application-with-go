package config

type appConfig struct {
	initialized  bool
	InProduction bool
}

var AppConfig appConfig

func init() {
	AppConfig = appConfig{
		initialized:  true,
		InProduction: false,
	}
}

func GetConfig() *appConfig {
	return &AppConfig
}
