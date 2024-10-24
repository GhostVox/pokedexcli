package types

import "github.com/Brent-the-carpenter/pokedexcli/internal/pokecache"

type Config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}
