package config

import (
	"log"
	"os"

	err "github.com/norbux/zet/pkg/err"
	"gopkg.in/yaml.v3"
)

// TODO: Handle database path in configuration
type Config struct {
	DatabaseName string `yaml:"database_name"`
}

func NewConfig() Config {

	file, e := os.ReadFile("/Users/norberto/dev/zet/config.yaml")
	err.Check(e)

	cfg := Config{}
	e = yaml.Unmarshal([]byte(file), &cfg)
	err.Check(e)

	log.Println("Read from the config file: ")
	log.Printf("%v", cfg)

	return Config{
		DatabaseName: cfg.DatabaseName,
	}
}
