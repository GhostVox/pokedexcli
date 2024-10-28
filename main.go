package main

import (
	"time"

	pokeapi "github.com/Brent-the-carpenter/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.PokemonInfo, 0),
	}

	startRepl(cfg)
}
