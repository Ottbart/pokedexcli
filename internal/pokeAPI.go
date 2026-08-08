package pokeAPI

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetPokeData(endpoint string) string {
	res, err := http.Get(endpoint)
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
	err := json.Unmarshal(jsonString, targetStruct)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
