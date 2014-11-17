package pollster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type DateEstimates []struct {
	Date      string `json:"date"`
	Estimates []struct {
		Choice string  `json:"choice"`
		Value  float32 `json:"value"`
	} `json:"estimates"`
}

type Estimates []struct {
	Choice         string  `json:"choice"`
	Value          float32 `json:"value"`
	LeadConfidence float32 `json:"lead_confidence"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Party          string  `json:"party"`
	Incumbent      bool    `json:"incumbent"`
}

type Chart struct {
	Title         string        `json:"title"`
	Slug          string        `json:"slug"`
	Topic         string        `json:"topic"`
	State         string        `json:"state"`
	ShortTitle    string        `json:"short_title"`
	ElectionDate  string        `json:"election_date"`
	PollCount     int           `json:"poll_count"`
	LastUpdated   time.Time     `json:"last_updated"`
	Url           string        `json:"url"`
	Estimates     Estimates     `json:"estimates"`
	DateEstimates DateEstimates `json:"estimates_by_date"`
}

type Poll struct {
	Id           int    `json:"id"`
	Pollster     string `json:"pollster"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Method       string `json:"method"`
	Source       string `json:"source"`
	LastUpdated  string `json:"last_updated"`
	Partisan     string `json:"partisan"`
	Affiliation  string `json:"affiliation"`
	SurveyHouses []struct {
		Name  string `json:"name"`
		Party string `json:"party"`
	} `json:"survey_houses"`
	Sponsors []struct {
		Name  string `json:"name"`
		Party string `json:"party"`
	}
	Questions []struct {
		Name           string `json:"name"`
		Chart          string `json:"chart"`
		Topic          string `json:"topic"`
		State          string `json:"state"`
		Subpopulations []struct {
			Name          string `json:"name"`
			Observations  int    `json:"observations"`
			MarginOfError int    `json:"margin_of_error"`
			Responses     []struct {
				Choice    string `json:"choice"`
				Value     int    `json:"value"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Party     string `json:"party"`
				Incumbent bool   `json:"incumbent"`
			} `json:"responses"`
		} `json:"subpopulations"`
	} `json:"questions"`
}

func handleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return false
}

func getJson(url string) []byte {
	res, err := http.Get(url)
	handleError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	handleError(err)
	return body
}

func buildUrl(url string, params map[string]string) string {
	for k, v := range params {
		url += fmt.Sprintf("%s=%s&", k, v)
	}
	return url
}

func Charts(params map[string]string) []Chart {
	url := buildUrl("http://elections.huffingtonpost.com/pollster/api/charts?", params)
	body := getJson(url)
	var charts []Chart
	json.Unmarshal(body, &charts)
	return charts
}

func (chart Chart) EstimatesByDate() DateEstimates {
	body := getJson(fmt.Sprintf("http://elections.huffingtonpost.com/pollster/api/charts/%s", chart.Slug))
	json.Unmarshal(body, &chart)
	return chart.DateEstimates
}

func Polls(params map[string]string) []Poll {
	url := buildUrl("http://elections.huffingtonpost.com/pollster/api/polls.json?", params)
	body := getJson(url)
	var polls []Poll
	json.Unmarshal(body, &polls)
	return polls
}
