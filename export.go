package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
	"github.com/internavenue/unlinked.in/schemas"
	"github.com/internavenue/unlinked.in/schemas/openhr"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ProfileExportHandler struct {
	store  sessions.Store
	config *Config
}

func (h *ProfileExportHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	state := req.FormValue("state")
	if len(code) == 0 || len(state) == 0 {
		http.Error(rw, "missing code or state param", http.StatusBadRequest)
		return
	}

	csrfToken, ok := getCSRFToken(h.store, req)
	if !ok || state != csrfToken {
		http.Error(rw, "state doesnot match CSRF-token", http.StatusForbidden)
		return
	}

	// get access token
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", h.config.SiteURL+"/auth/redirect")
	data.Set("client_id", h.config.APIKey)
	data.Set("client_secret", h.config.SecretKey)
	data.Set("code", code)

	resp, err := http.Post("https://www.linkedin.com/uas/oauth2/accessToken",
		"application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	if err != nil {
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}
	log.Info("accessToken response", resp.Status)

	// we don't care about expires_in
	token := struct {
		AccessToken string `json:"access_token"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&token)
	resp.Body.Close()
	if err != nil {
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}

	// get JSON data
	requestURL := "//api.linkedin.com/v1/people/~:" +
		"(email-address,id,first-name,last-name,maiden-name,formatted-name," +
		"phonetic-first-name,phonetic-last-name,formatted-phonetic-name," +
		"headline,location,industry,summary,specialties,positions,picture-url," +
		"site-standard-profile-request,public-profile-url,last-modified-timestamp," +
		"proposal-comments,associations,interests,skills,educations,courses," +
		"publications:(id,title,publisher:(name),authors:(id,name,person),date,url,summary)," +
		"patents:(id,title,summary,number,status,office,inventors:(id,name,person),date,url)," +
		"languages:(id,language:(name),proficiency:(level,name))," +
		"certifications:(id,name,authority:(name),number,start-date,end-date)," +
		"main-address,volunteer,three-current-positions,three-past-positions," +
		"num-recommenders,recommendations-received,date-of-birth," +
		"member-url-resources,phone-numbers,bound-account-types,im-accounts," +
		"twitter-accounts,primary-twitter-account)?format=json"

	r, _ := http.NewRequest("GET", "https://api.linkedin.com", nil)
	r.URL.Opaque = requestURL
	r.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		log.Println("Fetching data error", err)
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	profile := &schemas.LinkedInSchema{}
	err = json.Unmarshal(respBody, profile)
	if err != nil {
		log.Println("Decoding data error", err)
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}

	switch req.FormValue("format") {
	case "openhr":
		b, _ := json.MarshalIndent(openhr.FromLinkedInSchema(profile), "", "    ")
		rw.Write(b)
	default:
		rw.Write(respBody)
	}
}
