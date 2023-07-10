package main

import (
	"errors"
	"fmt"
)

func mapbCommand(cfg *config, area_name string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("You are on the first page")
	}

    locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
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
