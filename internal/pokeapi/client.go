package pokeapi

import (
	"net/http"
	"time"

	"github.com/ahgr3y/pokedex-repl-cli/internal/pokecache"
)

type Client struct {
	pokeCache  *pokecache.PokemonCache
	cache      *pokecache.Cache
	httpClient http.Client
}

func GetClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		pokeCache: pokecache.NewPokeCache(),
		cache:     pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
