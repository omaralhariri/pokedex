package main

import (
	"fmt"
)

func mapCommand(config *Config) error {
	if config.Next != "" && config.Previous == "" {
		config.Endpoint = config.Next
	}
	res := getPoke(config)
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	config.Next = res.Next
	config.Previous = res.Previous
	return nil
}
