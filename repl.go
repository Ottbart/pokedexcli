package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ottbart/pokedexcli/internal/pokeapi"
)

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
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
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 0 {
			args = cleanedInput[1:]
		}
		if cmd, exists := cfg.commands[command]; exists {
			if err := cmd.callback(cfg, args...); err != nil {
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

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			name:        "explore <area_name>",
			description: "Shows the list of Pokemon in the given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a Pokemon by name",
			callback:    commandCatch,
		},
	}
}
