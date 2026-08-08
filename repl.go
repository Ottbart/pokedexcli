package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeAPI "github.com/Ottbart/pokedexcli/internal"
)

type command struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Count    int    `json:"count"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cliCommands map[string]command

func startRepl() {
	cliCommands = map[string]command{
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display list of next areas on the map of the Pokemon world",
			callback:    commandMap,
		},
		"mapback": {
			name:        "mapb",
			description: "Display list of previous areas on the map of the Pokemon world",
			callback:    commandMapBack,
		},
	}
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}

	// Start an infinite loop to continuously read user input
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if cmd, exists := cliCommands[cleanedInput[0]]; exists {
			if err := cmd.callback(cfg); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

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
	// Display a list of areas on the map of the Pokemon world
	var endpoint string
	if cfg.Next == "" {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	} else {
		endpoint = cfg.Next
	}
	currentMap := pokeAPI.GetPokeData(endpoint)

	pokeAPI.JSON2Struct([]byte(currentMap), cfg)

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config) error {
	// Display a list of areas on the map of the Pokemon world
	var endpoint string
	if cfg.Previous == "" {
		endpoint = "https://pokeapi.co/api/v2/location-area/"
	} else {
		endpoint = cfg.Previous
	}
	currentMap := pokeAPI.GetPokeData(endpoint)

	pokeAPI.JSON2Struct([]byte(currentMap), cfg)

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func cleanInput(text string) []string {
	// Convert the input text to lowercase and split it into words, returning a slice of strings
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
