package main

import (
	"fmt"
	"net/http"
)

// ResirectHandler never returns an error and always returns http.StatusSeeOther
// with redirect header to the UI for handle authorization code or any error.
func ResirectHandler(s *Server, rw http.ResponseWriter, req *http.Request) (int, error) {
	errFormat := "/?error=%s&error_description=%s"
	if errMsg := req.FormValue("error"); len(errMsg) > 0 {
		errDesc := req.FormValue("error_description")
		http.Redirect(rw, req,
			fmt.Sprintf(errFormat, errMsg, errDesc), http.StatusSeeOther)
		return http.StatusSeeOther, nil
	}

	code := req.FormValue("code")
	state := req.FormValue("state")

	if len(code) == 0 || len(state) == 0 {
		http.Redirect(rw, req,
			fmt.Sprintf(errFormat, "missing_param", "missing code or state param"),
			http.StatusSeeOther)
		return http.StatusSeeOther, nil
	}

	csrfToken, ok := getCSRFToken(s.store, req)
	if !ok || state != csrfToken {
		http.Redirect(rw, req,
			fmt.Sprintf(errFormat, "invalid_state", "state doesnot match CSRF-token"),
			http.StatusSeeOther)
		return http.StatusSeeOther, nil
	}

	http.Redirect(rw, req,
		fmt.Sprintf("/?code=%s&state=%s", code, state), http.StatusSeeOther)
	return http.StatusSeeOther, nil
}
