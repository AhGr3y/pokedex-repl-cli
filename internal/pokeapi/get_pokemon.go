package pokeapi

import (
	"encoding/json"
	"io"

	"github.com/ahgr3y/pokedex-repl-cli/internal/poketype"
)

func (pokeapiClient *Client) GetPokemon(pokemon string) (poketype.Pokemon, error) {

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
		return poketype.Pokemon{}, err
	}
	defer resp.Body.Close()

	// Read data from resp
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return poketype.Pokemon{}, err
	}

	// Parse data to Pokemon struct
	pokemonSpecs := poketype.Pokemon{}
	err = json.Unmarshal(dat, &pokemonSpecs)
	if err != nil {
		return poketype.Pokemon{}, err
	}

	return pokemonSpecs, nil
}
