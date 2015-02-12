package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	var cfgfile string
	flag.StringVar(&cfgfile, "cfgfile", "./cfg.json", "path to JSON config file")

	config, err := ReadConfig(cfgfile)
	if err != nil {
		log.Fatal("Cannot load config from file", cfgfile, err)
	}

	tmpl := template.Must(template.
		ParseGlob(filepath.Join(config.TemplateDir, "*.tmpl")))
	store := sessions.NewCookieStore([]byte(config.SessionKey))

	r := mux.NewRouter()
	r.Handle("/auth/redirect", &Server{
		store:   store,
		config:  config,
		handler: ResirectHandler,
	})

	r.Handle("/api/export", &Server{
		store:   store,
		config:  config,
		handler: ProfileExportHandler,
	})

	r.Handle("/", &Server{
		store:    store,
		config:   config,
		template: tmpl,
		handler:  IndexHandler,
	})

	log.Info(http.ListenAndServe(config.HTTPAddress, r))
}
