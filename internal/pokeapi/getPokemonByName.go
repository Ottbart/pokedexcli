package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemonByName(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	var data []byte
	if cached, exists := c.cache.Get(url); exists {
		data = cached
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, data)
	}
	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
