package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// 1. CONFIG_ROOT
// 2. ${HOME}/.<app-name>
// 3. /etc/<app-name>
func init() {
	viper.SetConfigType("yaml")

	// 1. load CONFIG_ROOT
	if configRoot := os.Getenv("CONFIG_ROOT"); configRoot != "" {
		viper.AddConfigPath(configRoot)
	}

	if appName := os.Getenv("CONFIG_APP_NAME"); appName != "" {
		// 2. load from ${HOME}/.<app-name>
		if homeDir, err := os.UserHomeDir(); err == nil && homeDir != "" {
			viper.AddConfigPath(fmt.Sprintf("%v/.%v", homeDir, appName))
		}

		// 3. load from /etc/<app-name>
		viper.AddConfigPath(fmt.Sprintf("/etc/%v", appName))
	}

}
