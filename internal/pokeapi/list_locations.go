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
	locationType := location{}
	err = json.Unmarshal(dat, &locationType)
	if err != nil {
		return location{}, err
	}

	return locationType, nil

}
