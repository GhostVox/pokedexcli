package cli

import (
	"fmt"

	pokeapi "github.com/Brent-the-carpenter/pokedexcli/internal/pokeAPI"
)

type Config struct {
	next     *string
	previous *string
}
type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

func commandExit(config *Config) error {
	fmt.Println("Exiting...")
	return nil
}
func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commandMap := CliCommands()
	for key, value := range commandMap {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
func commandMap(config *Config) error {
	res, err := pokeapi.GetLocations(nil, config.next)
	if err != nil {
		fmt.Print("error occured while getting locations%w\n", err)
		return err
	}
	config.next = res.Next
	config.previous = res.Previous
	Results := res.Results
	for _, result := range Results {
		fmt.Println(result.Name)
	}

	return nil
}
func commandMapB(config *Config) error {
	res, err := pokeapi.GetLocations(nil, config.previous)
	if err != nil {
		fmt.Print("Error occurred while getting previous locations : %w\n", err)
		return err
	}
	config.next = res.Next
	config.previous = res.Previous
	results := res.Results
	for _, result := range results {
		println(result.Name)
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the pokemon world. If you are on the last page will return error.",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the pokemon world. If you are on the first page will return error.",
			Callback:    commandMapB,
		},
	}
}
