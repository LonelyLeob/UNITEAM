package internal

import (
	"fmt"
)

type ConfigDatabase struct {
	Port     string `env:"DB_PORT"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func (c *ConfigDatabase) GetFullDBAddr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Name)
}

type ConfigRemote struct {
	Port       string `env:"REMOTE_PORT"`
	SigningKey string `env:"S_KEY"`
}

func (c *ConfigRemote) GetFullWebAddr() string {
	return fmt.Sprintf(":%s", c.Port)
}

func (c *ConfigRemote) GetSKey() string {
	return fmt.Sprintf(c.SigningKey)
}
