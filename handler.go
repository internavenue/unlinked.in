package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

type HandlerError struct {
	Code        int
	Description string
	Detail      string
}

func (e *HandlerError) Error() string {
	return fmt.Sprintf("%s (code %d) %s", e.Description, e.Code, e.Detail)
}

type HandlerFunc func(*Server, http.ResponseWriter, *http.Request) (int, error)

type Server struct {
	store    sessions.Store
	config   *Config
	template *template.Template
	handler  HandlerFunc
}

func (h *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	status, err := h.handler(h, rw, req)
	if err != nil {
		rw.WriteHeader(status)
		errFields := logrus.Fields{"Status": status}
		e, ok := err.(*HandlerError)
		if ok {
			errFields["Code"] = e.Code
			if len(e.Description) > 0 {
				errFields["Description"] = e.Description
			}

			if len(e.Detail) > 0 {
				errFields["Detail"] = e.Detail
			}

			json.NewEncoder(rw).Encode(e)
		} else {
			json.NewEncoder(rw).
				Encode(&struct{ Description string }{err.Error()})
		}

		logrus.WithFields(errFields).Error("ServeHTTP error")
	}
}
