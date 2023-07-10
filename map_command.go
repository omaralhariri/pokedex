package main

import (
	"fmt"
)

func mapCommand(cfg *config, area_name string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, v := range locationsResp.Results {
		fmt.Println(v.Name)
	}

	return nil
}
