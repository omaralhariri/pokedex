package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(*Config) error
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

func startRepl() {
	config := Config{
		Endpoint: "https://pokeapi.co/api/v2/location",
		Next:     "",
		Previous: "",
	}

	for {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		word := input.Text()
		command, exists := getCommands()[word]
		if !exists {
			fmt.Println("Not a valid command")
			continue
		}
		command.callback(&config)
	}
}
