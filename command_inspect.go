package main

import "errors"

func commandInspect(cfg *config, args ...string) error {

	// Command should have one pokemon as argument
	if len(args) != 1 {
		return errors.New("please select one pokemon to inspect")
	}
	return nil
}
