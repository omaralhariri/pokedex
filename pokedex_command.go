package main

import (
	"fmt"
)

func pokedexCommand(cfg *config, area string) error {
	pokes, err := cfg.pokeapiClient.Pokedex()
	if err != nil {
		return err
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokes {
		fmt.Println("  - " + pokemon)
	}

	return nil
}
