package agentstats

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	ApiToken string
	client   *http.Client
}

type Group struct {
	Id   string `json:"groupid"`
	Name string `json:"groupname"`
	Rank string `json:"rank"`
}

func NewClient(apiToken string) *Client {
	return &Client{ApiToken: apiToken, client: &http.Client{}}
}

func (client *Client) Groups() (groups []Group, err error) {
	req, err := http.NewRequest("GET", "https://api.agent-stats.com/groups", nil)
	if err != nil {
		return
	}
	req.Header.Add("AS-Key", client.ApiToken)

	resp, err := client.client.Do(req)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&groups)

	return
}
