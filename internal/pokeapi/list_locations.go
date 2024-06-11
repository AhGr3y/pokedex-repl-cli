package pokeapi

import (
	"encoding/json"
	"io"

	"github.com/ahgr3y/pokedex-repl-cli/internal/poketype"
)

func (pokeapiClient *Client) ListLocations(resourceUrl *string) (poketype.Location, error) {

	url := baseUrl + "/location-area"
	if resourceUrl != nil {
		url = *resourceUrl
	}

	// If url exist in cache, use data in cache
	if dat, exist := pokeapiClient.cache.Get(url); exist {
		// Parse dat to location struct
		loc := poketype.Location{}
		err := json.Unmarshal(dat, &loc)
		if err != nil {
			return poketype.Location{}, err
		}
		return loc, nil
	}

	// Retrieve data from url
	resp, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return poketype.Location{}, err
	}
	defer resp.Body.Close()

	// Get data from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return poketype.Location{}, err
	}

	// Parse dat to location struct
	loc := poketype.Location{}
	err = json.Unmarshal(dat, &loc)
	if err != nil {
		return poketype.Location{}, err
	}

	// Save loc in cache
	pokeapiClient.cache.Add(url, dat)

	return loc, nil

}
