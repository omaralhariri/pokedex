package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/omaralhariri/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "Display a help message",
			callback: helpCommand,
		},
		"exit": {
			name:     "exit",
			desc:     "Exit the pokedex",
			callback: exitCommand,
		},
		"map": {
			name:     "map",
			desc:     "Display the names of 20 locations in the Pokemon world",
			callback: mapCommand,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Display the names of 20 previous locations in the Pokemon world",
			callback: mapbCommand,
		},
	}
}

func startRepl(cfg *config) {
	for {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		word := input.Text()
		command, exists := getCommands()[word]
		if !exists {
			fmt.Println("Not a valid command")
			continue
		} else {
            err := command.callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
        }
	}
}
