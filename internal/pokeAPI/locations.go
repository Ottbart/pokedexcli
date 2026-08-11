package pokeAPI

import (
	"encoding/json"
	"io"
)

func (c *Client) GetLocations(nextLocationsURL *string) (Locations, error) {
	url := baseURL + "/location-area/"
	if nextLocationsURL != nil {
		url = *nextLocationsURL
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Locations{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return Locations{}, err
	}

	locationsResp := Locations{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
