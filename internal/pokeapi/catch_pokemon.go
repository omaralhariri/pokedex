package pokeapi

import (
    "io"
    "net/http"
    "errors"
    "encoding/json"
    "math/rand"
)

func (c *Client)  CatchPokemon(pokename string) (string, error) {
    url := baseURL + "/pokemon/" + pokename

    if _, ok := c.cache.Get(pokename); ok {
        return pokename + " was caught!", nil
    }
    
    if val, ok := c.cache.Get(url); ok {
        respCatch := RespCatch{}
        err := json.Unmarshal(val, &respCatch)
        if err != nil {
            return "", err
        }

        if ! catchPoke(respCatch) {
            return "", errors.New(pokename + " escaped!")
        }

        c.cache.Add(pokename, val)
        return pokename + " was caught!", nil
    }

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

    c.cache.Add(url, body)

    if ! catchPoke(respCatch) {
        return "", errors.New(pokename + " escaped!")
    }

    c.cache.Add(pokename, body)

    return pokename + " was caught!", nil
}

func catchPoke(respCatch RespCatch) bool {

    chance := rand.Intn(100 + respCatch.BaseExperience)

    return chance > respCatch.BaseExperience
}
