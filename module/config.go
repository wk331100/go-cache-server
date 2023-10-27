package module

import (
	"path/filepath"

	"github.com/spf13/viper"
)

const configFilePath = "config/config.yml"

func ParseConfig() (*Config, error) {
	absPath, err := filepath.Abs(configFilePath)
	if err != nil {
		return nil, err
	}
	viper.SetConfigFile(absPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	Server     ServerConfig
	Persistent PersistentConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type PersistentConfig struct {
	Enable bool
	Mode   string
}
