package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/internavenue/unlinked.in/schemas"
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

	// we don't care about expires_in
	respBody := struct {
		AccessToken string `json:"access_token"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}

	// get JSON data
	requestURL := "https://api.linkedin.com/v1/people/~:"
	requestURL += "(id,first-name,last-name,headline,picture-url,industry,summary,specialties,positions:"
	requestURL += "(id,title,summary,start-date,end-date,is-current,company:"
	requestURL += "(id,name,type,size,industry,ticker)),educations:"
	requestURL += "(id,school-name,field-of-study,start-date,end-date,degree,activities,notes),"
	requestURL += "associations,interests,num-recommenders,date-of-birth,publications:"
	requestURL += "(id,title,publisher:(name),authors:(id,name),date,url,summary),"
	requestURL += "patents:(id,title,summary,number,status:(id,name),office:(name),"
	requestURL += "inventors:(id,name),date,url),languages:(id,language:(name),proficiency:(level,name)),"
	requestURL += "skills:(id,skill:(name)),certifications:(id,name,authority:"
	requestURL += "(name),number,start-date,end-date),courses:(id,name,number),"
	requestURL += "recommendations-received:(id,recommendation-type,recommendation-text,recommender),"
	requestURL += "honors-awards,three-current-positions,three-past-positions,volunteer)?format=json"

	r, _ := http.NewRequest("GET", requestURL, nil)
	r.Header.Set("Authorization", "Bearer "+respBody.AccessToken)
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}

	profile := &schemas.LinkedInProfile{}
	err = json.NewDecoder(resp.Body).Decode(profile)
	if err != nil {
		http.Error(rw, "Internal Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%#v", profile)
}
