package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocations(nextLocationsURL *string) (LocationRequest, error) {
	url := baseURL + "/location-area/"
	if nextLocationsURL != nil {
		url = *nextLocationsURL
	}
	var data []byte
	if cached, exists := c.cache.Get(url); exists {
		data = cached
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return LocationRequest{}, err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return LocationRequest{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationRequest{}, err
		}
		c.cache.Add(url, data)
	}
	locationsResp := LocationRequest{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationRequest{}, err
	}

	return locationsResp, nil
}
