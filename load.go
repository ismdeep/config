package config

import "github.com/spf13/viper"

// Load function
func Load(name string, data interface{}) error {
	viper.SetConfigName(name)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(data); err != nil {
		return err
	}

	return nil
}

// MustLoad function, will panic when meet errors
func MustLoad(name string, data interface{}) {
	if err := Load(name, data); err != nil {
		panic(err)
	}
}
