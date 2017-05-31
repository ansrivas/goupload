package main

import (
	"flag"
	"fmt"

	"github.com/ansrivas/goupload/logger"
)

var (
	config    string
	staticDir string
	log       = logger.Logger
)

func init() {

	flag.StringVar(&config, "config", "", "path to a configuration file")
	flag.StringVar(&staticDir, "staticDir", "", "path to a directory containing static files")
}

func main() {
	flag.Parse()

	if config == "" || staticDir == "" {
		log.Fatalln("Please check the config or staticDir")
	}
	fmt.Println(config, staticDir)
	fmt.Println("hello world")
}
