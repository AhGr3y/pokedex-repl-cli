package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
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
