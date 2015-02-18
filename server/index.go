package server

import (
	"net/http"
)

// IndexHandler executes index template. It generates set new CSRF if need.
// The function always returns (http.StatusOK, nil)
func IndexHandler(s *Server, rw http.ResponseWriter, req *http.Request) (int, error) {
	csrfToken, ok := getCSRFToken(s.store, req)
	if !ok {
		csrfToken = newCSRFToken(s.store, rw, req)
	}

	data := struct {
		CSRFToken string
		APIKey    string
		SiteURL   string
	}{csrfToken, s.config.APIKey, s.config.SiteURL}

	s.template.Execute(rw, &data)
	return http.StatusOK, nil
}
