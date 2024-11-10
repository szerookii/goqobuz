package qobuz

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/szerookii/goquobuz/qobuz/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	seedTimezoneRegex = regexp.MustCompile(`[a-z]\.initialSeed\("(?P<seed>[\w=]+)",window\.utimezone\.(?P<timezone>[a-z]+)\)`)
	infoExtrasRegex   = regexp.MustCompile(`name:"\w+/(?P<timezone>{timezones})",info:"(?P<info>[\w=]+)",extras:"(?P<extras>[\w=]+)"`)
	appIDRegex        = regexp.MustCompile(`production:{api:{appId:"(?P<app_id>\d{9})",appSecret:"(\w{32})`)
)

const (
	BaseURL    = "https://play.qobuz.com"
	APIBaseURL = "https://www.qobuz.com/api.json/0.2"
	LoginURL   = APIBaseURL + "/user/login"
)

type QobuzClient struct {
	email     string
	password  string
	user      *types.User
	authToken string
	loggedIn  bool

	app_id string
	secret string
}

type LoginReponse struct {
	User          *types.User `json:"user"`
	UserAuthToken string      `json:"user_auth_token"`
}

func NewFromCredentials(email, password string) (*QobuzClient, error) {
	c := &QobuzClient{
		email:    email,
		password: password,
	}

	_, _, err := c.GetAppIDAndSecrets()
	if err != nil {
		return nil, err
	}

	err = c.Login()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewFromAuthToken(authToken string) (*QobuzClient, error) {
	c := &QobuzClient{
		authToken: authToken,
	}

	_, _, err := c.GetAppIDAndSecrets()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *QobuzClient) Login() error {
	if s.loggedIn {
		return errors.New("already logged in")
	}

	params := url.Values{}
	params.Set("app_id", s.app_id)
	params.Set("email", s.email)
	params.Set("password", s.password)

	req, err := http.NewRequest("POST", LoginURL+"?"+params.Encode(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Qobuz/5.2.0 Android/29")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("login failed with status code %d", resp.StatusCode)
	}

	var loginResponse LoginReponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		return err
	}

	if loginResponse.User == nil || loginResponse.UserAuthToken == "" {
		return errors.New("login failed")
	}

	s.user = loginResponse.User
	s.authToken = loginResponse.UserAuthToken

	s.loggedIn = true

	return nil
}

func (s *QobuzClient) GetAppIDAndSecrets() (string, []string, error) {
	resp, err := http.Get("https://play.qobuz.com/login")
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	loginPage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	bundleURLRegex := regexp.MustCompile(`<script src="(/resources/\d+\.\d+\.\d+-[a-z]\d{3}/bundle\.js)"></script>`)
	bundleURLMatch := bundleURLRegex.FindSubmatch(loginPage)
	if bundleURLMatch == nil {
		return "", nil, fmt.Errorf("could not find bundle URL")
	}
	bundleURL := string(bundleURLMatch[1])

	resp, err = http.Get("https://play.qobuz.com" + bundleURL)
	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()
	bundleString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	bundle := string(bundleString)

	appIDMatch := appIDRegex.FindStringSubmatch(bundle)
	if appIDMatch == nil {
		return "", nil, fmt.Errorf("could not find app ID")
	}
	appID := appIDMatch[1]

	secrets := make(map[string][]string)
	seedMatches := seedTimezoneRegex.FindAllStringSubmatch(bundle, -1)
	for _, match := range seedMatches {
		seed, timezone := match[1], match[2]
		secrets[timezone] = []string{seed}
	}

	var orderedSecrets []string
	for timezone, _ := range secrets {
		orderedSecrets = append(orderedSecrets, timezone)
	}
	if len(orderedSecrets) > 1 {
		orderedSecrets[0], orderedSecrets[1] = orderedSecrets[1], orderedSecrets[0]
	}

	var capitalizedTimezones []string
	for _, timezone := range orderedSecrets {
		capitalizedTimezones = append(capitalizedTimezones, strings.Title(timezone))
	}

	infoExtrasRegex = regexp.MustCompile(strings.Replace(infoExtrasRegex.String(), "{timezones}", strings.Join(capitalizedTimezones, "|"), 1))

	infoExtrasMatches := infoExtrasRegex.FindAllStringSubmatch(bundle, -1)
	for _, match := range infoExtrasMatches {
		timezone, info, extras := strings.ToLower(match[1]), match[2], match[3]
		secrets[timezone] = append(secrets[timezone], info, extras)
	}

	var secretsList []string
	for _, timezone := range orderedSecrets {
		secret := strings.Join(secrets[timezone], "")
		if len(secret) > 44 {
			decodedSecret, err := base64.StdEncoding.DecodeString(secret[:len(secret)-44])
			if err == nil {
				secretsList = append(secretsList, string(decodedSecret))
			}
		}
	}

	s.app_id = appID
	s.secret = secretsList[len(secretsList)-1]

	return appID, secretsList, nil
}

type SearchResponse struct {
	Query  string `json:"query"`
	Albums struct {
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
		Total  int           `json:"total"`
		Items  []types.Album `json:"items"`
	} `json:"albums"`
	Tracks struct {
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
		Total  int           `json:"total"`
		Items  []types.Track `json:"items"`
	} `json:"tracks"`
	Artists struct {
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
		Total  int           `json:"total"`
		Items  []interface{} `json:"items"` // TODO: Implement artists
	} `json:"artists"`
	Playlists struct {
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
		Total  int           `json:"total"`
		Items  []interface{} `json:"items"` // TODO: Implement playlists
	} `json:"playlists"`
	Stories struct {
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
		Total  int           `json:"total"`
		Items  []interface{} `json:"items"` // TODO: Implement stories
	} `json:"stories"`
}

func (s *QobuzClient) Search(query string) (*SearchResponse, error) {
	req, err := http.NewRequest("GET", APIBaseURL+"/catalog/search?query="+url.QueryEscape(query), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-User-Auth-Token", s.authToken)
	req.Header.Set("X-App-Id", s.app_id)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var searchResponse SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return nil, err
	}

	return &searchResponse, nil
}
