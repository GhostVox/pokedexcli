package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Brent-the-carpenter/pokedexcli/internal/cli"
	"github.com/Brent-the-carpenter/pokedexcli/internal/pokecache"
	"github.com/Brent-the-carpenter/pokedexcli/types"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cli.CliCommands()
	startNext := "https://pokeapi.co/api/v2/location/"
	startPrevious := ""
	config := types.Config{
		Next:     &startNext,
		Previous: &startPrevious,
		Cache:    pokecache.NewCache(5 * time.Minute),
	}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			words := cleanInput(scanner.Text())
			input := string(words[0])
			if command, ok := commands[input]; ok {
				var err error
				if len(words) > 1 {
					err = command.Callback(&config, words[1])
				} else {
					err = command.Callback(&config, "")
				}
				if err != nil {
					fmt.Println("there was an error: %w", err)
				}
				fmt.Println()

			} else {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
			}
			if input == "exit" {
				break
			}
			fmt.Println()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("There was an error reading input: %v\n", err)
			break
		}

	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	splitInput := strings.Fields(input)
	return splitInput
}
