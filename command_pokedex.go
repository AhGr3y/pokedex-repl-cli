package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// This command should not have any argument
	if len(args) != 0 {
		return errors.New("this command does not take any argument, run 'pokedex' instead")
	}

	pokeCache := cfg.pokeCache

	// Send a message if user has not caught any pokemon
	if pokeCache.Len() == 0 {
		return errors.New("you have not caught any pokemon yet! go out there and catch em all")
	}

	fmt.Println("Your Pokedex:")
	userPokemon := pokeCache.GetPokemonList()
	for _, pokemon := range userPokemon {
		fmt.Printf("    - %v\n", pokemon)
	}

	return nil
}
