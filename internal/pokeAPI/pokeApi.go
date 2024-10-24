package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Brent-the-carpenter/pokedexcli/types"
)

type PokeResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type PokemonResponse struct {
	Region struct {
		Name string `json:"name"`
	} `json:"region"`
}

func GetLocations(testFunc func() (*http.Response, error), url *string, config *types.Config) (PokeResponse, error) {
	var res *http.Response
	var err error
	var data PokeResponse
	cache := config.Cache

	if testFunc != nil {
		res, err = testFunc()
	} else if cachedData, cached := cache.Get(*url); cached {
		err = json.Unmarshal(cachedData, &data)
		if err != nil {
			return PokeResponse{}, fmt.Errorf("error unmarshalling cached data: %w", err)
		}
		return data, nil
	} else {
		res, err = http.Get(*url)
	}
	if err != nil {
		return PokeResponse{}, fmt.Errorf("error making request to pokeAPI error:%w", err)
	}

	response, err := io.ReadAll(res.Body)
	config.Cache.Add(*url, response)
	if err != nil {
		return PokeResponse{}, fmt.Errorf("error reading response: %w", err)
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return PokeResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return data, nil
}

// Gets area  from second word passes into command
func GetPokemon(config *types.Config, area string) (PokemonResponse, error) {
	if len(area) == 0 {
		return PokemonResponse{}, fmt.Errorf("ldocation undefined")
	}
	baseUrl := "https://pokeapi.co/api/v2/location/"
	var data PokemonResponse
	if res, cached := config.Cache.Get(baseUrl + area); cached {
		json.Unmarshal(res, &data)
		return data, nil
	}
	res, err := http.Get(baseUrl + area)
	if err != nil {
		return PokemonResponse{}, err
	}

	parsedRes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error occurred while reading data from response: %w", err)
		return PokemonResponse{}, err
	}
	config.Cache.Add(baseUrl+area, parsedRes)

	err = json.Unmarshal(parsedRes, &data)
	if err != nil {
		fmt.Println("Error occurred while unmarshilling data into structered response.")
		return PokemonResponse{}, err
	}
	return data, nil

}
