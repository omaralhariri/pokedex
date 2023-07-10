package main

import (
	"fmt"
	"strconv"
)

func inspectCommand(cfg *config, pokename string) error {
	respCatch, err := cfg.pokeapiClient.InspectPokemon(pokename)

	if err != nil {
		return err
	}

	fmt.Println("Name: " + pokename)
	fmt.Println("Height: " + strconv.Itoa(respCatch.Height))
	fmt.Println("Weight: " + strconv.Itoa(respCatch.Weight))

	fmt.Println("Stats:")
	for _, stat := range respCatch.Stats {
		fmt.Println("   -" + stat.Stat.Name + ": " + strconv.Itoa(stat.BaseStat))
	}

	fmt.Println("Types:")
	for _, ty := range respCatch.Types {
		fmt.Println("  - " + ty.Type.Name)
	}

	return nil
}
