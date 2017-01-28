package gograph

import (
	"encoding/json"
	"net/http"
	"strings"
)

const currentVersion = "2.8"
const basePath = "https://graph.facebook.com"

// FacebookReq is a facebookReq
type FacebookReq struct {
	accessToken string
	APIVersion  string
}

// FacebookResponse is a thing
type FacebookResponse struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// New is a constructor
func New(accessToken string) FacebookReq {

	return FacebookReq{
		accessToken: accessToken,
	}
}

func (fb *FacebookReq) request(method, node string, fields []string) (*FacebookResponse, error) {
	req, err := http.NewRequest(method, basePath+"/"+fb.APIVersion+"/"+node, nil)
	profile := new(FacebookResponse)

	q := req.URL.Query()
	q.Add("access_token", fb.accessToken)
	req.URL.RawQuery = q.Encode()

	f := strings.Join(fields, ",")
	q.Add("fields", f)

	res, err := http.Get(req.URL.String())
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(profile)

	return profile, err
}

func (fb *FacebookReq) get(node string, fields []string) (*FacebookResponse, error) {
	return fb.request("GET", node, fields)
}

// Me fetches profile of current user
func (fb *FacebookReq) Me(fields []string) (*FacebookResponse, error) {
	return fb.get("me", fields)
}
