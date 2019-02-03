package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RemoteOKJob struct {
	Slug        string   `json:"slug"`
	ID          string   `json:"id"`
	Epoch       string   `json:"epoch"`
	Date        string   `json:"date"`
	Company     string   `json:"company"`
	CompanyLogo string   `json:"company_logo"`
	Position    string   `json:"position"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
}

func GetRemoteOKJobs() []RemoteOKJob {
	r, err := http.Get("https://remoteok.io/api?tags=golang")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	data := make([]RemoteOKJob, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	return data
}
