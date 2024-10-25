package cli

import (
	"fmt"

	pokeapi "github.com/Brent-the-carpenter/pokedexcli/internal/pokeAPI"
	"github.com/Brent-the-carpenter/pokedexcli/types"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*types.Config, string) error
}

func commandExit(config *types.Config, _ string) error {
	config.Cache.Stop()
	fmt.Println("Exiting...")
	return nil
}
func commandHelp(config *types.Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commandMap := CliCommands()
	for key, value := range commandMap {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
func commandMap(config *types.Config, _ string) error {
	res, err := pokeapi.GetLocations(nil, config.Next, config)
	if err != nil {
		fmt.Print("error occured while getting locations%w\n", err)
		return err
	}
	config.Next = res.Next
	config.Previous = res.Previous
	Results := res.Results
	for _, result := range Results {
		fmt.Println(result.Name)
	}

	return nil
}
func commandMapB(config *types.Config, _ string) error {
	if config.Previous == nil {
		return fmt.Errorf("can not go back, you at the starting point")
	}
	res, err := pokeapi.GetLocations(nil, config.Previous, config)
	if err != nil {
		fmt.Print("Error occurred while getting previous locations : %w\n", err)
		return err
	}
	config.Next = res.Next
	config.Previous = res.Previous
	results := res.Results
	for _, result := range results {
		println(result.Name)
	}
	return nil
}
func commandExplore(config *types.Config, area string) error {

	if area == "" {
		return fmt.Errorf("invalid or missing location")
	}
	fmt.Printf("Exploring %v...\n", area)
	res, err := pokeapi.GetPokemon(config, area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, element := range res.Encounters {
		fmt.Printf("- %v\n", element.Pokemon.Name)
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
		"explore": {
			name:        "explore",
			description: "Takes name of area as argument and returns all of the pokemon in the area.",
			Callback:    commandExplore,
		},
	}
}
