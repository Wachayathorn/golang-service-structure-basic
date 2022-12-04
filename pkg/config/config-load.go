package config

import (
	"fmt"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type ConfigInterface interface {
	Load(env string) error
}

type Config struct {
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
}

type State string

const (
	StateLocal State = "local"
	StateDev   State = "dev"
)

func (c *Config) ParseState(env string) State {
	switch env {
	case "dev":
		return StateDev
	default:
		return StateLocal
	}
}

func (c *Config) Load(logger echo.Logger, env string) (string, error) {
	vn := viper.New()
	vn.AddConfigPath(filepath.Join(".", "pkg", "config", "resource"))     // Set path to load config.yml
	vn.SetConfigName(fmt.Sprintf("config.%s", c.ParseState(string(env)))) // Load config file

	if err := vn.ReadInConfig(); err != nil { // Read config from file
		logger.Fatalf("Read config error:%s", err.Error())
		return "", err
	}

	if err := c.binding(vn); err != nil {
		logger.Fatalf("Binding config error:%s", err.Error())
		return "", err
	}

	logger.Infof("Start server with site:%s", c.ParseState(string(env)))
	return c.Port, nil
}

func (c *Config) binding(vn *viper.Viper) error {
	if err := vn.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}
