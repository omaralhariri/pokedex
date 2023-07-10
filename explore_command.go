package main

import "fmt"

func exploreCommand(cfg *config, area_name string) error {
	exploreArea, err := cfg.pokeapiClient.ExploreLocation(area_name)

	if err != nil {
		return err
	}

	fmt.Println("Exploring " + area_name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range exploreArea.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
