package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Brent-the-carpenter/pokedexcli/cli"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cli.CliCommands()
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()

			if command, ok := commands[input]; ok {

				err := command.Callback(commands)
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
