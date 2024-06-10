package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	// Display error if user did not provide location area
	if len(args) != 1 {
		return errors.New("please provide one location area to explore")
	}

	locationArea := args[0]

	fmt.Printf("Exploring %v...\n", locationArea)
	fmt.Println("Found Pokemon:")

	// Retrieve encounter data from pokeapi
	enc, err := cfg.pokeapiClient.ListEncounters(locationArea)
	if err != nil {
		return err
	}

	// Loop through pokemon encounters and print each pokemon
	pokeEncounters := enc.PokemonEncounters
	for _, enc := range pokeEncounters {
		fmt.Printf("- %v\n", enc.Pokemon.Name)
	}

	return nil
}
