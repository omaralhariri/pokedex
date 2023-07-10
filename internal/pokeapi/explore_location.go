package pokeapi

import (
    "io"
    "encoding/json"
    "net/http"
)

func (c *Client) ExploreLocation(area_name string) (RespExplore, error) {
    url := baseURL + "/location-area/" + area_name

    if val, ok := c.cache.Get(url); ok {
        exploreArea := RespExplore{}
        err := json.Unmarshal(val, &exploreArea)

        if err != nil {
            return RespExplore{}, err
        }

        return exploreArea, nil
    }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
       return RespExplore{}, err
	}

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespExplore{}, err
    }

	body, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespExplore{}, err
    }

    exploreArea := RespExplore{}
    err = json.Unmarshal(body, &exploreArea)
    if err != nil {
        return RespExplore{}, err
    }

    c.cache.Add(url, body)

    return exploreArea, nil
}
