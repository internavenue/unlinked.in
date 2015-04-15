package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/internavenue/unlinked.in/server"
	"net/http"
)

func main() {
	var cfgfile string
	flag.StringVar(&cfgfile, "cfgfile", "./cfg.json", "path to JSON config file")
	flag.Parse()

	config, err := server.ReadConfig(cfgfile)
	if err != nil {
		log.Fatal("Cannot load config from file", cfgfile, err)
	}

	log.Info(http.ListenAndServe(config.HTTPAddress, server.NewHandler(config)))
}
