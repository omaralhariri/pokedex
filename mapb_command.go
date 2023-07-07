package main

import (
	"errors"
	"fmt"
)

func mapbCommand(config *Config) error {
	if config.Previous == "" {
		fmt.Println("No previous pages are available")
		return errors.New("No previous pages are available")
	}

	config.Endpoint = config.Previous
	res := getPoke(config)
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	config.Next = res.Next
	config.Previous = res.Previous

	return nil
}
