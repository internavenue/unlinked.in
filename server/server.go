package server

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"path/filepath"
)

func NewHandler(config *Config) *mux.Router {
	tmpl := template.Must(template.
		ParseGlob(filepath.Join(config.TemplateDir, "*.tmpl")))
	store := sessions.NewCookieStore([]byte(config.SessionKey))

	r := mux.NewRouter()
	r.Handle("/auth/redirect", &Server{
		store:   store,
		config:  config,
		handler: ResirectHandler,
	})

	r.Handle("/export", &Server{
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

	return r
}
