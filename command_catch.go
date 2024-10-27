package main

func commandCatch(cfg *config, pokemonName string) error {
	pokemon, err := cfg.pokeapiClient.GetpokemonInfo()
	if err != nil {
		return err
	}

	return nil
}
