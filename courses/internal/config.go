package internal

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigRemote struct {
	BindPort   string `env:"REMOTE_PORT"`
	SigningKey string `env:"S_KEY"`
}

type ConfigDatabase struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func SetConfigs() (*ConfigRemote, *ConfigDatabase) {
	var cr ConfigRemote
	var cd ConfigDatabase

	if err := cleanenv.ReadEnv(&cr); err != nil {
		log.Fatalf("config: %s", err.Error())
	}

	if err := cleanenv.ReadEnv(&cd); err != nil {
		log.Fatalf("config: %s", err.Error())
	}

	return &cr, &cd
}

func (cr *ConfigRemote) CreateDomainAddr() string {
	return fmt.Sprintf(":%s", cr.BindPort)
}

func (c *ConfigDatabase) CreateFullDBAddr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Name)
}
