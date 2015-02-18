package server

import (
	"encoding/json"
	"github.com/internavenue/unlinked.in/schemas"
	"github.com/internavenue/unlinked.in/schemas/europass"
	"github.com/internavenue/unlinked.in/schemas/jsonresume"
	"github.com/internavenue/unlinked.in/schemas/openhr"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func ProfileExportHandler(s *Server,
	rw http.ResponseWriter, req *http.Request) (int, error) {
	code := req.FormValue("code")
	state := req.FormValue("state")
	if len(code) == 0 || len(state) == 0 {
		return http.StatusBadRequest, &HandlerError{
			Description: "missing code or state param",
		}
	}

	csrfToken, ok := getCSRFToken(s.store, req)
	if !ok || state != csrfToken {
		return http.StatusForbidden, &HandlerError{
			Description: "state doesnot match CSRF-token",
		}
	}

	// get access token
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", s.config.SiteURL+"/auth/redirect")
	data.Set("client_id", s.config.APIKey)
	data.Set("client_secret", s.config.SecretKey)
	data.Set("code", code)

	resp, err := http.Post("https://www.linkedin.com/uas/oauth2/accessToken",
		"application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	if err != nil {
		return http.StatusInternalServerError, &HandlerError{
			Description: err.Error(),
			Detail:      "Error when request access token",
		}
	}

	if resp.StatusCode != http.StatusOK {
		return http.StatusInternalServerError, &HandlerError{
			Description: resp.Status,
			Detail:      "Error response returned when request access token",
		}
	}

	// we don't care about expires_in
	token := struct {
		AccessToken string `json:"access_token"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&token)
	resp.Body.Close()
	if err != nil {
		return http.StatusInternalServerError, &HandlerError{
			Description: err.Error(),
			Detail:      "Error when parsing message returned from request access token",
		}
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
		"num-recommenders,recommendations-received,date-of-birth,honors-awards," +
		"member-url-resources,phone-numbers,bound-account-types,im-accounts," +
		"twitter-accounts,primary-twitter-account)?format=json"

	r, _ := http.NewRequest("GET", "https://api.linkedin.com", nil)
	r.URL.Opaque = requestURL
	r.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		return http.StatusInternalServerError, &HandlerError{
			Description: err.Error(),
			Detail:      "Error when request user data",
		}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	profile := &schemas.LinkedInSchema{}
	err = json.Unmarshal(respBody, profile)
	if err != nil {
		return http.StatusInternalServerError, &HandlerError{
			Description: err.Error(),
			Detail:      "Error when parse data returned from request user data",
		}
	}

	switch req.FormValue("format") {
	case "openhr":
		b, _ := json.MarshalIndent(openhr.FromLinkedInSchema(profile), "", "  ")
		rw.Write(b)
	case "jsonresume":
		b, _ := json.MarshalIndent(jsonresume.FromLinkedInSchema(profile), "", "  ")
		rw.Write(b)
	case "europass":
		b, _ := json.MarshalIndent(europass.FromLinkedInSchema(profile), "", "  ")
		rw.Write(b)
	default:
		rw.Write(respBody)
	}

	return http.StatusOK, nil
}
