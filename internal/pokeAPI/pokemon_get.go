package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonInfo, error) {
	fullURL := baseURL + "pokemon/" + name

	req, err := http.NewRequest("get", fullURL, nil)
	if err != nil {
		return PokemonInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	var pokemonResp PokemonInfo
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonInfo{}, err
	}

	return pokemonResp, nil
}
