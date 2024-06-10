package pokeapi

import (
	"encoding/json"
	"io"
)

func (pokeapiClient *Client) ListLocations(resourceUrl *string) (location, error) {

	url := baseUrl + "/location-area"
	if resourceUrl != nil {
		url = *resourceUrl
	}

	// If url exist in cache, use data in cache
	if dat, exist := pokeapiClient.cache.Get(url); exist {
		// Parse dat to location struct
		loc := location{}
		err := json.Unmarshal(dat, &loc)
		if err != nil {
			return location{}, err
		}
		return loc, nil
	}

	// Retrieve data from url
	resp, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return location{}, err
	}
	defer resp.Body.Close()

	// Get data from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return location{}, err
	}

	// Parse dat to location struct
	loc := location{}
	err = json.Unmarshal(dat, &loc)
	if err != nil {
		return location{}, err
	}

	// Save loc in cache
	pokeapiClient.cache.Add(url, dat)

	return loc, nil

}
