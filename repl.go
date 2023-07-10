package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"

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
	callback func(*config, string) error
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
		"explore": {
			name:     "explore",
			desc:     "Display the names of pokemons in a given area",
			callback: exploreCommand,
		},
	}
}

func startRepl(cfg *config) {
	for {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		words := input.Text()
        words_slice := strings.Split(words, " ")
        word := words_slice[0]
		command, exists := getCommands()[word]
		if !exists {
			fmt.Println("Not a valid command")
			continue
		} else {
            var area string
            if word == "explore" {
                if len(words_slice) < 2 {
                    fmt.Println("You have to provide an area to explore")
                    continue
                }
                area = words_slice[1]
            }

            err := command.callback(cfg, area)
            if err != nil {
                fmt.Println(err)
            }
        }
	}
}
