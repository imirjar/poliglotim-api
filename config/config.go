package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConf  `yaml:"server" envPrefix:"SERVER_"`
	Database StorageConf `yaml:"database" envPrefix:"DB_"`
}

type ServerConf struct {
	Port    string `yaml:"port" env:"PORT"`
	PubKey  string `yaml:"pub_key" env:"PUBKEY"`
	PrivKey string `yaml:"priv_key" env:"PRIVKEY"`
}

type StorageConf struct {
	Name string `yaml:"name" env:"NAME"`
	User string `yaml:"user" env:"USER"`
	Pswd string `yaml:"password" env:"PASSWORD"`
	Host string `yaml:"host" env:"HOST"`
	Port string `yaml:"port" env:"PORT"`
}

func (db *StorageConf) GetConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Pswd, db.Host, db.Port, db.Name)
}

func New() *Config {

	cfg := Config{}
	err := cfg.readFile("config/config.yml")
	if err != nil {
		log.Print(err)
		return nil
	}
	err = cfg.readEnv()
	if err != nil {
		log.Print(err)
		return nil
	}
	return &cfg
}

func (cfg *Config) readFile(configPath string) error {
	if configPath != "" {
		data, err := os.ReadFile(configPath)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return err
		}
		log.Println(cfg.Database.Name, cfg.Server.Port)
		return nil
	}
	return fmt.Errorf("no config file path")
}

func (cfg *Config) readEnv() error {
	err := env.Parse(cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return err

}
