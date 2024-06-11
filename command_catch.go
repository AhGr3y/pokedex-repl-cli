package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {

	// Only one pokemon can be caught at a time
	if len(args) != 1 {
		return errors.New("please choose one pokemon to catch")
	}

	// Catch pokemon
	pokemon := args[0]
	caught, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	if caught {
		fmt.Printf("Yes! %v was caught!\n", pokemon)
		return nil
	} else {
		fmt.Printf("Darn it! %v escaped!\n", pokemon)
	}

	return nil
}
