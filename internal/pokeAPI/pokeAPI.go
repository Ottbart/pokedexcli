package pokeAPI

import (
	"encoding/json"
	"io"
	"log"
)

func (c *Client) GetPokeData(endpoint string) string {
	// Use the httpClient from the Client struct to make the GET request
	res, err := c.httpClient.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func JSON2Struct(jsonString []byte, targetStruct interface{}) error {
	// Unmarshal the JSON string into the provided target struct
	err := json.Unmarshal(jsonString, targetStruct)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
