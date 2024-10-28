package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(" you must provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, caught := cfg.caughtPokemon[pokemonName]
	if !caught {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name:%s\n", pokemonName)
	fmt.Printf("Height:%d\n", pokemon.Height)
	fmt.Printf("Weight:%d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}
