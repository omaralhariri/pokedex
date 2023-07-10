package pokeapi

import (
    "strings"
    "errors"
)

func (c *Client) Pokedex() ([]string, error) {
    if val, ok := c.cache.Get("caught_pokes"); ok {
        names := string(val)
        return strings.Split(names, " "), nil
    } else {
        return []string{}, errors.New("You have no caught pokemons")
    }
}

