package internal

import (
	"errors"
	"fmt"

	"github.com/ansrivas/logger"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	//ErrConfigUninitalized is thrown whenever a configuration is not initialized.
	ErrConfigUninitalized = fmt.Errorf("Config uninitialized. Please call InitConfig first")
	log                   = logger.Logger
)

// InitConfig returns the config object for this project
func InitConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {

		return errors.New("Config file not found: " + configPath)
	}
	config = viper.GetViper()
	return nil
}

//GetConfig returns the viper config object
func GetConfig() (*viper.Viper, error) {
	if config != nil {
		return config, nil
	}
	return nil, ErrConfigUninitalized
}
