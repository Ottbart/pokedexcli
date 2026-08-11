package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ottbart/pokedexcli/internal/pokeapi"
)

type command struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

var cliCommands map[string]command

func startRepl(cfg *config) {
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
		"mapb": {
			name:        "mapb",
			description: "Display list of previous areas on the map of the Pokemon world",
			callback:    commandMapBack,
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
			if err := cmd.callback(cfg); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Unknown command: %s\nType 'help' for a list of available commands.\n", cleanedInput[0])
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func cleanInput(text string) []string {
	// Convert the input text to lowercase and split it into words, returning a slice of strings
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
