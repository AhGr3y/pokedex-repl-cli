package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	// Command should have one pokemon as argument
	if len(args) != 1 {
		return errors.New("please select one pokemon to inspect")
	}

	pokemon := args[0]

	// Cannot inspect a pokemon that has not been caught
	pokemonSpecs, exist := cfg.pokeCache.Get(pokemon)
	if !exist {
		return errors.New("unable to inspect pokemon that has not been caught yet")
	}

	// Print out details of pokemon
	fmt.Printf("Name: %v\n", pokemonSpecs.Name)
	fmt.Printf("Height: %v\n", pokemonSpecs.Height)
	fmt.Printf("Weight: %v\n", pokemonSpecs.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonSpecs.Stats {
		fmt.Printf("    - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonSpecs.Types {
		fmt.Printf("    - %v\n", t.Type.Name)
	}

	return nil
}
