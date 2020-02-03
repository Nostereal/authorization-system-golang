package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nostereal/login-system/internal/app/apiserver"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to your config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	lvl, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Printf("Error while parsing log_level: %v", err)
	}
	logrus.SetLevel(lvl)

	if err := apiserver.Start(config); err != nil {
		logrus.Fatal(err)
	}
}
