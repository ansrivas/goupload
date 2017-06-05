package internal

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var (
	//ErrConfigUninitalized is thrown whenever a configuration is not initialized.
	ErrConfigUninitalized = fmt.Errorf("Config uninitialized. Please call InitConfig first")
)

//Configuration holds a config object from viper.
type Configuration struct {
	config *viper.Viper
}

// NewConfig returns the config object for this project
func NewConfig(configPath string) (*Configuration, error) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {

		return nil, errors.New("Config file not found: " + configPath)
	}
	return &Configuration{config: viper.GetViper()}, nil
}

//GetConfig returns the viper config object
func (config Configuration) GetConfig() *viper.Viper {
	return config.config
}
