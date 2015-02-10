package main

import (
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

type IndexHandler struct {
	store    sessions.Store
	config   *Config
	template *template.Template
}

func (h *IndexHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	csrfToken, ok := getCSRFToken(h.store, req)
	if !ok {
		csrfToken = newCSRFToken(h.store, rw, req)
	}

	data := struct {
		CSRFToken string
		APIKey    string
		SiteURL   string
	}{csrfToken, h.config.APIKey, h.config.SiteURL}

	h.template.Execute(rw, &data)
}
