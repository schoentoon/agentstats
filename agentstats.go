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

type Progress struct {
	AP             uint64 `json:"ap"`
	Explorer       uint64 `json:"explorer"`
	Seer           uint64 `json:"seer"`
	Collector      uint64 `json:"collector"`
	Hacker         uint64 `json:"hacker"`
	Builder        uint64 `json:"builder"`
	Connector      uint64 `json:"connector"`
	MindController uint64 `json:"mind-controller"`
	MU             uint64 `json:"illuminator"`
	Binder         uint16 `json:"binder"`
	CountryMaster  uint64 `json:"country-master"`
	Recharger      uint64 `json:"recharger"`
	Liberator      uint64 `json:"liberator"`
	Pioneer        uint64 `json:"pioneer"`
	Purifier       uint64 `json:"purifier"`
	Neutralizer    uint64 `json:"neutralizer"`
	Disruptor      uint64 `json:"disruptor"`
	Salvator       uint64 `json:"salvator"`
	Trekker        uint64 `json:"trekker"`
	Guardian       uint64 `json:"guardian"`
	Smuggler       uint64 `json:"smuggler"`
	LinkMaster     uint64 `json:"link-master"`
	Controller     uint64 `json:"controller"`
	FieldMaster    uint64 `json:"field-master"`
	SpecOps        uint64 `json:"specops"`
	Engineer       uint64 `json:"engineer"`
	Sojourner      uint64 `json:"sojourner"`
	Recruiter      uint64 `json:"recruiter"`
	Translator     uint64 `json:"translator"`
	MissionDay     uint64 `json:"missionday"`
	Level          uint8  `json:"level"`
	Faction        string `json:"faction"`
	LastSubmit     string `json:"last_submit"`
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

func (client *Client) Progress(group string) (progress map[string]Progress, err error) {
	req, err := http.NewRequest("GET", "https://api.agent-stats.com/groups/"+group, nil)
	if err != nil {
		return
	}
	req.Header.Add("AS-Key", client.ApiToken)

	resp, err := client.client.Do(req)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&progress)

	return
}
