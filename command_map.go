package main

import (
	"errors"
	"fmt"
)

var isLastPage = false
var isFirstPage = true

func commandMapf(cfg *config, args ...string) error {
	// This command should not have any arguments
	if len(args) != 0 {
		return errors.New("this command does not take any argument, run 'map' instead")
	}

	// Print error message if user reaches last page
	if isLastPage {
		return errors.New("unable to go to the next page, this is the last page")
	}

	// Fetch location data from url
	location, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	// Udate config with new next and prev location url
	cfg.nextLocationURL = location.Next
	cfg.prevLocationURL = location.Previous

	// Display location areas
	locationAreas := location.Results
	for i := range locationAreas {
		fmt.Println(locationAreas[i].Name)
	}

	// Set isLastPage to true when user reaches last page
	if location.Next == nil {
		isLastPage = true
	}

	// Set isFirstPage to false after user leave first page
	if location.Previous != nil {
		isFirstPage = false
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	// This command should not have any arguments
	if len(args) != 0 {
		return errors.New("this command does not take any argument, run 'mapb' instead")
	}

	// Print error message if user reaches first page
	if isFirstPage {
		return errors.New("unable to go to the previous page, this is the first page")
	}

	// Fetch location data from url
	location, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	// Udate config with new next and prev location url
	cfg.nextLocationURL = location.Next
	cfg.prevLocationURL = location.Previous

	// Display location areas
	locationAreas := location.Results
	for i := range locationAreas {
		fmt.Println(locationAreas[i].Name)
	}

	// Set isFirstPage to true when user reaches first page
	if location.Previous == nil {
		isFirstPage = true
	}

	// Set isLastPage to false since there is always a next page once we go to a previous page
	if location.Next != nil {
		isLastPage = false
	}

	return nil
}
