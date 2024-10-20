package cli

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(map[string]cliCommand) error
}

func commandExit(commandMap map[string]cliCommand) error {
	fmt.Println("Exiting...")
	return nil
}
func commandHelp(commandMap map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for key, value := range commandMap {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func CliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			Callback:    commandExit,
		},
	}
}
