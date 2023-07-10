package main

import "fmt"

func catchCommand(cfg *config, pokename string) error {
	respCatch, err := cfg.pokeapiClient.CatchPokemon(pokename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(respCatch)
	return nil
}
