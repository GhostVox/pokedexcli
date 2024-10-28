package main

import (
	"fmt"
)

func commandpokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you have not caught any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for key := range cfg.caughtPokemon {
		fmt.Printf("- %s", key)
	}
	return nil
}
