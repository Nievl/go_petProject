package main

import (
	"flag"
	"log"
	"petProject/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.yml", "path to config file")
}

func main() {
	flag.Parse()
	config := server.NewConfig()
	err := config.ParseConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
