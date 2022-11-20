package internal

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigRemote struct {
	BindPort string `env:"RPORT"`
}

type ConfigMongo struct {
	Host     string `env:"MHOST"`
	Port     string `env:"MPORT"`
	User     string `env:"MUSER"`
	Password string `env:"MPWD"`
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
