package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getPoke(config *Config) Response {
	res, err := http.Get(config.Endpoint)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("Error reading: ", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println("Error not nil: ", err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
	}

	return response
}
