package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	// Exit the program with a status code of 0 (success)
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	// Display available commands and their descriptions
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	// Display a list of next areas on the map of the Pokemon world
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	// Display a list of previous areas on the map of the Pokemon world
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cfg *config, areaName ...string) error {
	// Display a list of Pokemon in the current area
	if len(areaName) != 1 {
		return errors.New("please provide a valid area name to explore")
	}
	Location, err := cfg.pokeapiClient.GetPokemonInArea(areaName[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", Location.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range Location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, pokemonName ...string) error {
	// Catch a Pokemon by name
	if len(pokemonName) != 1 {
		return errors.New("please provide a valid Pokemon name to catch")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemonByName(pokemonName[0])
	if err != nil {
		return err
	}
	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("Caught %s!\n", pokemon.Name)
	return nil
}
