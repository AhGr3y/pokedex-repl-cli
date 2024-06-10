package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ahgr3y/pokedex-repl-cli/internal/pokeapi"
)

var cliName = "Pokedex"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

// Used for mapping commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Displays a list of pokemon encountered in the specified location area",
			callback:    commandExplore,
		},
	}
}

// Normalize user input
func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	words := strings.Fields(lowerCase)
	return words
}

// Start the Pokedex
func startRepl(cfg *config) {

	// Display welcome message to get user started
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Enter 'help' to get the list of available commands.")

	// Process user commands until they give the 'exit' command
	for {
		// Display prompt for user to enter commands
		fmt.Print(cliName + " > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		commands := cleanInput(scanner.Text())

		// Prompt again if user provides empty command
		if len(commands) == 0 {
			continue
		}

		args := []string{}
		if len(commands) > 1 {
			args = commands[1:]
		}
		commandName := commands[0]
		command, exist := getCommands()[commandName]

		if exist {

			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}

		} else {
			fmt.Println("Invalid command")
			continue
		}
	}
}
