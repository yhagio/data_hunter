package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GithubJob struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyLogo string `json:"company_logo"`
	CompanyURL  string `json:"company_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Location    string `json:"location"`
	HowToApply  string `json:"how_to_apply"`
}

func GetGithubJobs() []GithubJob {
	r, err := http.Get("https://jobs.github.com/positions.json?search=golang&page=0")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	data := make([]GithubJob, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	return data
}
