package pokeapi

import (
    "errors"
    "encoding/json"
)

func (c *Client) InspectPokemon(pokename string) (RespCatch, error) {
    if val, ok := c.cache.Get(pokename); ok {
        respCatch := RespCatch{}
        err := json.Unmarshal(val, &respCatch)
        if err != nil {
            return RespCatch{}, err
        }

        return respCatch, nil
    }

    return RespCatch{}, errors.New("you have not caught that pokemon")
}
