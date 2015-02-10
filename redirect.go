package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
	"net/http"
)

type AuthRedirectHandler struct {
	store  sessions.Store
	config *Config
}

func (h *AuthRedirectHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	errFormat := "/?error=%s&error_description=%s"
	if errMsg := req.FormValue("error"); len(errMsg) > 0 {
		errDesc := req.FormValue("error_description")
		http.Redirect(rw, req,
			fmt.Sprintf(errFormat, errMsg, errDesc), http.StatusSeeOther)
		return
	}

	code := req.FormValue("code")
	state := req.FormValue("state")

	if len(code) == 0 || len(state) == 0 {
		http.Redirect(rw, req,
			fmt.Sprintf(errFormat, "missing_param", "missing code or state param"),
			http.StatusSeeOther)
		return
	}

	csrfToken, ok := getCSRFToken(h.store, req)
	if !ok || state != csrfToken {
		http.Redirect(rw, req,
			fmt.Sprintf(errFormat, "invalid_state", "state doesnot match CSRF-token"),
			http.StatusSeeOther)
		return
	}

	http.Redirect(rw, req,
		fmt.Sprintf("/?code=%s&state=%s", code, state), http.StatusSeeOther)
}
