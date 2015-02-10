package main

import (
	"encoding/base64"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

func getCSRFToken(store sessions.Store, req *http.Request) (string, bool) {
	session, _ := store.Get(req, "csrf-token")
	token, ok := session.Values["token"].(string)
	if len(token) != 40 {
		return "", false
	}

	return token, ok
}

func newCSRFToken(store sessions.Store, rw http.ResponseWriter, req *http.Request) string {
	session, _ := store.Get(req, "csrf-token")
	token := base64.URLEncoding.EncodeToString(securecookie.GenerateRandomKey(30))
	session.Values["token"] = token
	return token
}
