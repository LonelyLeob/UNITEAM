package forms

import (
	"fmt"
)

type ConfigDatabase struct {
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	Name     string `env:"NAME"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

func (c *ConfigDatabase) GetFullDBAddr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Name)
}

type ConfigRemote struct {
	Port       string `env:"REM_PORTF"`
	Host       string `env:"REM_HOSTF"`
	SigningKey string `env:"S_KEY"`
}

func (c *ConfigRemote) GetFullWebAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *ConfigRemote) GetSKey() string {
	return fmt.Sprintf(c.SigningKey)
}
