package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")

	// load user define config root via CONFIG_ENV_NAME environment
	if envName := os.Getenv("CONFIG_ENV_NAME"); envName != "" {
		if envValue := os.Getenv(envName); envValue != "" {
			viper.AddConfigPath(envValue)
		}
	}

	// load CONFIG_ROOT
	if configRoot := os.Getenv("CONFIG_ROOT"); configRoot != "" {
		viper.AddConfigPath(configRoot)
	}

	// load from /etc/<app-name>
	if appName := os.Getenv("CONFIG_APP_NAME"); appName != "" {
		viper.AddConfigPath(fmt.Sprintf("/etc/%v", appName))
	}

}
