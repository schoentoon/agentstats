package agentstats

import (
	"encoding/json"
	"net/http"
	"time"
)

// Client An api client used to make calls to the agent stats api
type Client struct {
	APIToken string
	client   *http.Client
}

// Group A representation of a group, you can later on use the ID to request stats of people in the group
type Group struct {
	ID   string `json:"groupid"`
	Name string `json:"groupname"`
	Rank string `json:"rank"`
}

// Stats A representation of stats of an agent, you'll get these from the GroupProgress() call
type Stats struct {
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

// Progress Top level object from the Progress() call
type Progress struct {
	Medals Medals `json:"mymedals"`
}

// Medals All the medals that we have data for in the Progress struct
type Medals struct {
	Explorer       Medal `json:"explorer"`
	Seer           Medal `json:"seer"`
	Collector      Medal `json:"collector"`
	Hacker         Medal `json:"hacker"`
	Builder        Medal `json:"builder"`
	Connector      Medal `json:"connector"`
	MindController Medal `json:"mind-controller"`
	MU             Medal `json:"illuminator"`
	Binder         Medal `json:"binder"`
	CountryMaster  Medal `json:"country-master"`
	Recharger      Medal `json:"recharger"`
	Liberator      Medal `json:"liberator"`
	Pioneer        Medal `json:"pioneer"`
	Purifier       Medal `json:"purifier"`
	Neutralizer    Medal `json:"neutralizer"`
	Disruptor      Medal `json:"disruptor"`
	Salvator       Medal `json:"salvator"`
	Trekker        Medal `json:"trekker"`
	Guardian       Medal `json:"guardian"`
	Smuggler       Medal `json:"smuggler"`
	LinkMaster     Medal `json:"link-master"`
	Controller     Medal `json:"controller"`
	FieldMaster    Medal `json:"field-master"`
	SpecOps        Medal `json:"specops"`
	Engineer       Medal `json:"engineer"`
	Sojourner      Medal `json:"sojourner"`
	Recruiter      Medal `json:"recruiter"`
	Translator     Medal `json:"translator"`
	MissionDay     Medal `json:"missionday"`
}

// Medal A single medal and that info about it, like how many are missing for black etc.
type Medal struct {
	Missing struct {
		Black    uint64 `json:"black"`
		Platinum uint64 `json:"platinum"`
		Gold     uint64 `json:"gold"`
		Silver   uint64 `json:"silver"`
		Bronze   uint64 `json:"bronze"`
	} `json:"miss"`
	Progress struct {
		Latest uint64 `json:"latest"`
		Week   uint64 `json:"week"`
		Month  uint64 `json:"month"`
		Total  uint64 `json:"total"`
	} `json:"progression"`
}

// NewClient Create a new client using the api token provided
func NewClient(APIToken string) *Client {
	return &Client{APIToken: APIToken, client: &http.Client{}}
}

// Groups List all the groups that the authenticated user is in
func (client *Client) Groups() (groups []Group, err error) {
	req, err := http.NewRequest("GET", "https://api.agent-stats.com/groups", nil)
	if err != nil {
		return
	}
	req.Header.Add("AS-Key", client.APIToken)

	resp, err := client.client.Do(req)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&groups)

	return
}

// GroupProgress Show all the progress of all the people in a certain group
func (client *Client) GroupProgress(group string) (progress map[string]Stats, err error) {
	req, err := http.NewRequest("GET", "https://api.agent-stats.com/groups/"+group, nil)
	if err != nil {
		return
	}
	req.Header.Add("AS-Key", client.APIToken)

	resp, err := client.client.Do(req)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&progress)

	return
}

// Progress Get detailed progress information about the authenticated user
func (client *Client) Progress(since time.Time) (output Progress, err error) {
	req, err := http.NewRequest("GET", "https://api.agent-stats.com/progress/"+since.Format("2006-01-02"), nil)
	if err != nil {
		return
	}
	req.Header.Add("AS-Key", client.APIToken)

	resp, err := client.client.Do(req)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&output)

	return
}
