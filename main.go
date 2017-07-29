package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ansrivas/goupload/internal"
	"github.com/ansrivas/goupload/logger"
	"github.com/ansrivas/goupload/routes"
	"github.com/spf13/viper"
)

var (
	configPath string
)

func init() {

	flag.StringVar(&configPath, "configPath", "", "path to a configuration file")
}

func bootstrap(configPath string) (*viper.Viper, error) {
	config, err := internal.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	return config.GetConfig(), nil
}

func main() {
	flag.Parse()
	log := logger.SetupLogger("goupload")

	if configPath == "" {
		log.Fatalln("Please check the config ")
	}

	config, err := bootstrap(configPath)
	if err != nil {
		log.Fatalln("Unable to read configurations file")
	}
	assertDir := config.GetString("app.static_dir_path")
	templateList := internal.SetAssetsPath(assertDir)

	router, err := routes.NewRouter(templateList, log, config)

	if err != nil {
		log.Errorln(err.Error())
	}
	// config and init server
	addr := config.GetString("app.host_port")
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Info(fmt.Sprintf("server listening on %s", addr))
	log.Fatal(s.ListenAndServe())
}
