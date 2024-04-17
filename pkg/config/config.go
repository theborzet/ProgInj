package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Host     string
	DBPort   string
	User     string
	Password string
	DBname   string
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	dbConfig := Config{
		Port:     viper.GetString("PORT"),
		Host:     viper.GetString("DB_HOST"),
		DBPort:   viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		DBname:   viper.GetString("DB_NAME"),
	}

	return &dbConfig, nil
}

func (c *Config) Connection() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Host, c.DBPort, c.User, c.Password, c.DBname)
}
