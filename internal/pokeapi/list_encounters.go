package pokeapi

import (
	"encoding/json"
	"io"
)

func (pokeapiClient *Client) ListEncounters(location string) (encounter, error) {
	url := baseUrl + "/location-area/" + location

	// If url exist in cache, use data in cache
	if dat, exist := pokeapiClient.cache.Get(url); exist {
		// Parse dat to location struct
		enc := encounter{}
		err := json.Unmarshal(dat, &enc)
		if err != nil {
			return encounter{}, err
		}
		return enc, nil
	}

	// Retrieve data from url
	resp, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return encounter{}, err
	}
	defer resp.Body.Close()

	// Get data from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return encounter{}, err
	}

	// Parse dat to location struct
	enc := encounter{}
	err = json.Unmarshal(dat, &enc)
	if err != nil {
		return encounter{}, err
	}

	// Save loc in cache
	pokeapiClient.cache.Add(url, dat)

	return enc, nil
}
