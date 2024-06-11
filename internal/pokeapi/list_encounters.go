package pokeapi

import (
	"encoding/json"
	"io"

	"github.com/ahgr3y/pokedex-repl-cli/internal/poketype"
)

func (pokeapiClient *Client) ListEncounters(location string) (poketype.Encounter, error) {
	url := baseUrl + "/location-area/" + location

	// If url exist in cache, use data in cache
	if dat, exist := pokeapiClient.cache.Get(url); exist {
		// Parse dat to location struct
		enc := poketype.Encounter{}
		err := json.Unmarshal(dat, &enc)
		if err != nil {
			return poketype.Encounter{}, err
		}
		return enc, nil
	}

	// Retrieve data from url
	resp, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return poketype.Encounter{}, err
	}
	defer resp.Body.Close()

	// Get data from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return poketype.Encounter{}, err
	}

	// Parse dat to location struct
	enc := poketype.Encounter{}
	err = json.Unmarshal(dat, &enc)
	if err != nil {
		return poketype.Encounter{}, err
	}

	// Save loc in cache
	pokeapiClient.cache.Add(url, dat)

	return enc, nil
}
