package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Env *Config

type WorkNode struct {
	Number  int    `json:"number"`
	SymId   string `json:"symId"`
	HttpUrl string `json:"httpUrl"`
	WsUrl   string `json:"wsUrl"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type HostConfig struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type CacheConfig struct {
	Interval uint `json:"interval"`
	Use      bool `json:"use"`
}

type BootNodeConfig struct {
	HttpUrl string `json:"httpUrl"`
}

type Config struct {
	Host      HostConfig        `json:"host"`
	Domain    string            `json:"domain"`
	Cache     CacheConfig       `json:"cache"`
	WorkNodes []*WorkNode       `json:"workNodes"`
	BootNodes []*BootNodeConfig `json:"bootNodes"`
	Database  DatabaseConfig    `json:"database"`
}

func LoadEnvConfig(filename string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Env = config
	return config
}

func GetEnv() *Config {
	return Env
}
