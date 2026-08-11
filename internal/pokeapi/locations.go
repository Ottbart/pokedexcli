package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocations(nextLocationsURL *string) (Locations, error) {
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
			return Locations{}, err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return Locations{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Locations{}, err
		}
		c.cache.Add(url, data)
	}
	locationsResp := Locations{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
