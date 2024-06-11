package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	// This command should not have any arguments
	if len(args) != 0 {
		return errors.New("this command does not take any argument, run 'help' instead")
	}

	commands := getCommands()
	fmt.Println("========================================")
	fmt.Println("Here are the list of available commands:")
	fmt.Println()
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	fmt.Println("========================================")
	return nil
}
