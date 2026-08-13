package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Ottbart/pokedexcli/internal/pokeapi"
)

// Pseudocode plan:
// 1) Build a test helper that captures stdout because commandPokedex prints to the terminal.
// 2) Test the empty-pokedex case: cfg.caughtPokemon is empty, run commandPokedex, and verify the "You haven't caught any Pokemon yet." text appears.
// 3) Test the populated-pokedex case: cfg.caughtPokemon contains at least two Pokemon, run commandPokedex, and verify the header and names are printed.
// 4) Fail fast on unexpected errors from commandPokedex while preserving the output capture logic.

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() failed: %v", err)
	}

	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("w.Close() failed: %v", err)
	}

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("io.ReadAll() failed: %v", err)
	}

	return string(out)
}

func TestCommandPokedex_NoPokemon(t *testing.T) {
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	output := captureStdout(t, func() {
		if err := commandPokedex(cfg); err != nil {
			t.Fatalf("commandPokedex returned unexpected error: %v", err)
		}
	})

	if !strings.Contains(output, "You haven't caught any Pokemon yet.") {
		t.Fatalf("expected empty pokedex message, got: %q", output)
	}
}

func TestCommandPokedex_WithPokemon(t *testing.T) {
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{
			"pikachu":   {Name: "pikachu"},
			"bulbasaur": {Name: "bulbasaur"},
		},
	}

	output := captureStdout(t, func() {
		if err := commandPokedex(cfg); err != nil {
			t.Fatalf("commandPokedex returned unexpected error: %v", err)
		}
	})

	if !strings.Contains(output, "Caught Pokemon:") {
		t.Fatalf("expected pokedex header, got: %q", output)
	}
	if !strings.Contains(output, "pikachu") {
		t.Fatalf("expected pikachu in output, got: %q", output)
	}
	if !strings.Contains(output, "bulbasaur") {
		t.Fatalf("expected bulbasaur in output, got: %q", output)
	}
}
