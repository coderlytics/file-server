package server

import (
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Configuration is the configuration structure which holds all config values
type Configuration struct {
	Server     Server     `mapstructure:"server"`
	FileServer FileServer `mapstructure:"file-server"`
	Logging    Logging    `mapstructure:"logging"`
}

// Server holds the basic configuration for the server
type Server struct {
	Port string `mapstructure:"port" validate:"number" default:"5140"`
}

// FileServer holds the basic configuration for the served files
type FileServer struct {
	Files []File
}

// Logging configures the logging sub system
type Logging struct {
	LogLevel string `mapstructure:"level" validate:"oneof=trace debug info warn error fatal panic" default:"error"`
}

// File is the definition of a single servable file
type File struct {
	Name     string `mapstructure:"name"`
	FilePath string `mapstructure:"file"`
	Endpoint string `mapstructure:"endpoint"`
}

var Config Configuration

// initConfig sets up the configuration system by reading the configuration file from the given path
func initConfig(conf string) error {
	clear()

	if err := setDefaultValues(); err != nil {
		return err
	}

	if err := read(conf); err != nil {
		return err
	}

	return verify()

}

// read is reading the values from the configuration file and putting it in the Configuration struct
func read(configFile string) error {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Config)

	if err != nil {
		return err
	}

	return nil
}

// setDefaultValues writes the default values of the configuration to the Configuration struct
func setDefaultValues() error {
	return defaults.Set(&Config)
}

// verify checks the configuration file for semantic errors
func verify() error {
	validate := validator.New()

	return validate.Struct(Config)
}

// Clear removes the whole configuration.
// To use the configuration Init must be called again
func clear() {
	Config = Configuration{}
}
