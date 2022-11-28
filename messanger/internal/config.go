package internal

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigRemote struct {
	BindPort string `env:"REMOTE_PORT"`
}

type ConfigMongo struct {
	Host     string `env:"MONGO_HOST"`
	Port     string `env:"MONGO_PORT"`
	User     string `env:"MONGO_USER"`
	Password string `env:"MONGO_PWD"`
}

func SetConfigs() (*ConfigRemote, *ConfigMongo) {
	var cr ConfigRemote
	var cm ConfigMongo

	if err := cleanenv.ReadEnv(&cr); err != nil {
		log.Fatalf("config: %s", err.Error())
	}

	if err := cleanenv.ReadEnv(&cm); err != nil {
		log.Fatalf("config: %s", err.Error())
	}

	return &cr, &cm
}

func (cr *ConfigRemote) CreateDomainAddr() string {
	return fmt.Sprintf(":%s", cr.BindPort)
}

func (cm *ConfigMongo) CreateURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", cm.User, cm.Password, cm.Host, cm.Port)
}
