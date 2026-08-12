package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemonInArea(locationAreaName string) (Location, error) {
	url := baseURL + "/location-area/" + locationAreaName
	var data []byte
	if cached, exists := c.cache.Get(url); exists {
		data = cached
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return Location{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}
		c.cache.Add(url, data)
	}
	locationResp := Location{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	return locationResp, nil
}
