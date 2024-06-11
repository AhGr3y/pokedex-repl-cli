package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	// Only one pokemon can be caught at a time
	if len(args) != 1 {
		return errors.New("please choose one pokemon to catch")
	}

	// Catch pokemon
	pokemon := args[0]
	pokemonSpecs, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	// Calculate difficulty of catching pokemon using pokemon's base experience
	catchDifficulty := getCatchDifficulty(pokemonSpecs.BaseExperience)
	// Get user's rolled int value; 0 to 99
	userRolledInt := rand.Intn(100)

	if userRolledInt <= catchDifficulty { // if user caught successfully

		// Add pokemon to pokeCache
		cfg.pokeCache.Add(pokemon, pokemonSpecs)

		fmt.Printf("Yes! %v was caught!\n", pokemon)
	} else {
		fmt.Printf("Darn it! %v escaped!\n", pokemon)
	}

	return nil
}

func getCatchDifficulty(baseExp int) int {
	if baseExp > 400 {
		return 10
	}
	if baseExp > 350 {
		return 20
	}
	if baseExp > 300 {
		return 30
	}
	if baseExp > 250 {
		return 40
	}
	if baseExp > 200 {
		return 50
	}
	if baseExp > 150 {
		return 60
	}
	if baseExp > 100 {
		return 70
	}
	if baseExp > 50 {
		return 80
	}
	return 90
}
