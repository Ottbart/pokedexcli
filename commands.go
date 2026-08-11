package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	// Exit the program with a status code of 0 (success)
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	// Display available commands and their descriptions
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cliCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
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

func commandMapBack(cfg *config) error {
	// Display a list of previous areas on the map of the Pokemon world
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
