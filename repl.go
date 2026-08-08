package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ottbart/pokedexcli/internal/pokeAPI"
)

type command struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	previous string
	next     string
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
			description: "Display list of areas on the map of the Pokemon world",
			callback:    commandMap,
		},
	}
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Start an infinite loop to continuously read user input
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if cmd, exists := cliCommands[cleanedInput[0]]; exists {
			cmd.callback()
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
	type Map struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}
	mapData := Map{}
	endpoint := "location-area/" //empty location area endpoint to get bunch of location areas
	currentMap := pokeAPI.GetPokeData(endpoint)
	pokeAPI.json2struct([]byte(currentMap), &mapData)
	return nil
}

func cleanInput(text string) []string {
	// Convert the input text to lowercase and split it into words, returning a slice of strings
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
