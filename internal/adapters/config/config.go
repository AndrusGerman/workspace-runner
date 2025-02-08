package config

import (
	"os"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"

	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig() *Config {

	godotenv.Load()
	return &Config{}
}

func (c *Config) Host() string {
	return ":" + c.Port()
}

func (c *Config) Debug() bool {
	return true
}

func (c *Config) Port() string {
	return c.getValue("PORT", "4500")

}

func (c *Config) Database() types.Database {
	return "workspace-runnerDB"
}

func (c *Config) UriMongo() string {
	return "mongodb://user:passwod@localhost:27017/"
}

func (c *Config) APP_NAME() string {
	return c.getValue("APP_NAME", "github.com/AndrusGerman/workspace-runner")
}

func (c *Config) getValue(envName string, defaultValue string) string {
	if os.Getenv(envName) != "" {
		return os.Getenv(envName)
	}
	return defaultValue
}
