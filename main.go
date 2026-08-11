package main

import (
	"time"

	"github.com/Ottbart/pokedexcli/internal/pokeAPI"
)

func main() {
	pokeClient := pokeAPI.NewClient(5*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
