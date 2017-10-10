package gograph

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const currentVersion = "2.9"
const basePath = "https://graph.facebook.com"

const CouldNotPerformRequest = "Could not perform request"

// FacebookReq is a facebookReq
type FacebookReq struct {
	accessToken string
	APIVersion  string
}

// FacebookResponse is a thing
type FacebookResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        string `json:"id"`
}

// New is a constructor

func New(accessToken string) FacebookReq {
	return FacebookReq{
		APIVersion:  currentVersion,
		accessToken: accessToken,
	}
}

func (fb *FacebookReq) request(method, node string, fields []string) (*FacebookResponse, error) {
	req, err := http.NewRequest(method, basePath+"/v"+fb.APIVersion+"/"+node, nil)
	profile := new(FacebookResponse)

	q := req.URL.Query()
	q.Add("access_token", fb.accessToken)
	f := strings.Join(fields, ",")
	q.Add("fields", f)
	req.URL.RawQuery = q.Encode()

	res, err := http.Get(req.URL.String())

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		json.NewDecoder(res.Body).Decode(profile)
		return profile, nil
	}

	return nil, errors.New(CouldNotPerformRequest)
}

func (fb *FacebookReq) get(node string, fields []string) (*FacebookResponse, error) {
	return fb.request("GET", node, fields)
}

// Me fetches profile of current user
func (fb *FacebookReq) Me(fields []string) (*FacebookResponse, error) {
	return fb.get("me", fields)
}
