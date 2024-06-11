package pokeapi

import (
	"encoding/json"
	"io"
	"math/rand"

	"github.com/ahgr3y/pokedex-repl-cli/internal/poketype"
)

func (pokeapiClient *Client) CatchPokemon(pokemon string) (bool, error) {

	/*
		Attempt to catch pokemon
		If fail to catch, return false
		If succeed to catch, return true and add to PokemonCache
	*/

	// Get resource url
	url := baseUrl + "/pokemon/" + pokemon

	// Fetch data from API
	resp, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Read data from resp
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Parse data to Pokemon struct
	pokemonSpecs := poketype.Pokemon{}
	err = json.Unmarshal(dat, &pokemonSpecs)
	if err != nil {
		return false, err
	}

	// Calculate difficulty of catching pokemon using pokemon's base experience
	catchDifficulty := getCatchDifficulty(pokemonSpecs.BaseExperience)
	// Get user's rolled int value; 0 to 99
	userRolledInt := rand.Intn(100)

	if userRolledInt <= catchDifficulty { // if user caught successfully

		// Add pokemon to pokemon cache
		pokeapiClient.pokeCache.Add(pokemon, pokemonSpecs)

		return true, nil
	}

	return false, nil
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
