package anilist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	clientId     string
	clientSecret string
	accessToken  string
	expires      time.Time
}

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int64  `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
}

// IsExpired returns true if the session expired.
func (s *Session) IsExpired() bool {
	return s.accessToken == "" || s.expires.IsZero() || time.Now().After(s.expires)
}

// Auth retrieves a new access token.
func (s *Session) Auth() (err error) {
	var (
		authURL *url.URL

		req  *http.Request
		resp *http.Response

		rawData []byte
	)

	authURL, err = APIURL.Parse("auth/access_token")
	if err != nil {
		return
	}

	query := authURL.Query()
	query.Set("grant_type", "client_credentials")
	query.Set("client_id", s.clientId)
	query.Set("client_secret", s.clientSecret)

	authURL.RawQuery = query.Encode()

	req, err = http.NewRequest("POST", authURL.String(), nil)
	if err != nil {
		return
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	rawData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ar := &authResponse{}

	err = json.Unmarshal(rawData, &ar)
	if err != nil {
		return
	}

	s.accessToken = ar.AccessToken
	s.expires = time.Unix(ar.Expires, 0)

	if s.IsExpired() {
		err = errors.New("could not authenticate")
		return
	}

	return
}

func (s *Session) NewRequest(method, urlStr string, body io.Reader) (req *http.Request, err error) {
	var (
		reqURL *url.URL
	)

	if s.IsExpired() {
		err = errors.New("session expired")
		return
	}

	reqURL, err = url.Parse(urlStr)
	if err != nil {
		return
	}

	query := reqURL.Query()
	query.Set("access_token", s.accessToken)

	reqURL.RawQuery = query.Encode()

	req, err = http.NewRequest(method, reqURL.String(), body)
	if err != nil {
		return
	}

	req.Header.Set("X-TOKEN", s.accessToken)

	return
}

func (s *Session) escapeTerms(terms string) (escapedTerms string) {
	escapedTerms = url.QueryEscape(terms)
	escapedTerms = strings.Replace(escapedTerms, "+", "%20", -1)
	return
}

func (s *Session) searchPages(searchURL *url.URL) (pages [][]byte, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)

	pages = make([][]byte, 0)

	query := searchURL.Query()

	for page := 1; ; page++ {
		var (
			rawBody []byte
		)

		query.Set("page", strconv.Itoa(page))
		searchURL.RawQuery = query.Encode()

		req, err = s.NewRequest("GET", searchURL.String(), nil)
		if err != nil {
			return
		}

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		rawBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		rawBodyStr := string(rawBody)
		rawBodyStr = strings.TrimSpace(rawBodyStr)
		if rawBodyStr == "" || rawBodyStr == "[]" {
			break
		}

		pageResults := make([]interface{}, 0)
		err = json.Unmarshal(rawBody, &pageResults)
		if err != nil {
			return
		}

		if len(pageResults) == 0 {
			break
		}

		pages = append(pages, rawBody)
	}

	return
}

func (s *Session) SearchAnime(terms string) (results []*AnimeSmall, err error) {
	var (
		searchURL *url.URL

		pages [][]byte
	)

	escapedTerms := s.escapeTerms(terms)
	searchURLStr := fmt.Sprintf("anime/search/%s", escapedTerms)

	searchURL, err = APIURL.Parse(searchURLStr)
	if err != nil {
		return
	}

	pages, err = s.searchPages(searchURL)
	if err != nil {
		return
	}

	results = make([]*AnimeSmall, 0)

	for _, rawBody := range pages {
		pageResults := make([]*AnimeSmall, 0)
		err = json.Unmarshal(rawBody, &pageResults)
		if err != nil {
			return
		}

		if len(pageResults) == 0 {
			break
		}

		results = append(results, pageResults...)
	}

	return
}

func (s *Session) SearchCharacter(terms string) (results []*CharacterSmall, err error) {
	var (
		searchURL *url.URL

		pages [][]byte
	)

	escapedTerms := s.escapeTerms(terms)
	searchURLStr := fmt.Sprintf("character/search/%s", escapedTerms)

	searchURL, err = APIURL.Parse(searchURLStr)
	if err != nil {
		return
	}

	pages, err = s.searchPages(searchURL)
	if err != nil {
		return
	}

	results = make([]*CharacterSmall, 0)

	for _, rawBody := range pages {

		pageResults := make([]*CharacterSmall, 0)
		err = json.Unmarshal(rawBody, &pageResults)
		if err != nil {
			return
		}

		if len(pageResults) == 0 {
			break
		}

		results = append(results, pageResults...)
	}

	return
}

func NewSession(clientId, clientSecret string) (s *Session, err error) {
	if clientId == "" {
		err = errors.New("empty clientId")
		return
	}
	if clientSecret == "" {
		err = errors.New("empty clientSecret")
		return
	}

	s = &Session{
		clientId:     clientId,
		clientSecret: clientSecret,
	}

	return
}
