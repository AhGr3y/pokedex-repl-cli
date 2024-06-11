package pokeapi

import (
	"net/http"
	"time"

	"github.com/ahgr3y/pokedex-repl-cli/internal/pokecache"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func GetClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
