package main

import (
	"errors"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	// This command should not have any arguments
	if len(args) != 0 {
		return errors.New("this command does not take any argument, run 'exit' instead")
	}

	os.Exit(0)
	return nil
}
