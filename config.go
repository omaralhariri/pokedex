package main

type Config struct {
	Endpoint string
	Next     string
	Previous string
}

func getConfig() Config {
	return Config{
		Endpoint: "https://pokeapi.co/api/v2/location/",
		Next:     "https://pokeapi.co/api/v2/location/",
		Previous: "https://pokeapi.co/api/v2/location",
	}
}
