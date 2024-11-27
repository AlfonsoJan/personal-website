package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `yaml:"database"`
}

var AppConfig *Config

func SetConfigFile() error {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("can't find the file config.yaml :%s", err)
	}

	env := viper.GetString("environment")
	if env == "" {
		return fmt.Errorf("environment not specified in config.yaml")
	}

	envConfig := viper.Sub(env)
	if envConfig == nil {
		return fmt.Errorf("configuration for environment '%s' not found in config.yaml", env)
	}

	if err := envConfig.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("unable to load configuration for environment '%s': %v", env, err)
	}

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("no .env file found or unable to read")
	}

	if dbPassword, ok := viper.Get("DB_PASSWORD").(string); ok && dbPassword != "" {
		AppConfig.Database.Password = dbPassword
	}
	return nil
}
