package pokeapi

import (
    "io"
    "fmt"
    "net/http"
    "errors"
    "encoding/json"
    "math/rand"
)

func (c *Client)  CatchPokemon(pokename string) (string, error) {
    url := baseURL + "/pokemon/" + pokename

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    respCatch := RespCatch{}
    err = json.Unmarshal(body, &respCatch)
    if err != nil {
        return "", err
    }

    chance := rand.Intn(100 + respCatch.BaseExperience)

    if respCatch.BaseExperience > chance {
        return "", errors.New(pokename + " escaped!")
    }

    c.cache.Add(pokename, body)

    return pokename + " was caught!", nil
}

