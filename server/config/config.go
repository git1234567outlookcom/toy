package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type tomlConfig struct {
	Title string
	DB    database  `toml:"database"`
	Log   logconfig `toml:"logconfig"`
}

type logconfig struct {
	Level string
}
type database struct {
	Username string
	Password string
	Uri      string
	Name     string
}

var Config tomlConfig

func InitConfig(args []string) {
	postfix := "config_local.toml"
	if len(args) > 1 {
		postfix = args[1]
		switch postfix {
		case "test":
			postfix = "config_test.toml"
		case "prod":
			postfix = "config_prod.toml"
		default:
			postfix = "config_local.toml"
		}
	}
	path, _ := os.Getwd()
	log.Println("Path:\t", path)
	if _, err := toml.DecodeFile("./config/"+postfix, &Config); err != nil {
		log.Println(err)
		return
	}
	log.Printf("Config:\t %+v\n", Config)
}
