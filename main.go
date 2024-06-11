package main

import (
	"time"

	"github.com/ahgr3y/pokedex-repl-cli/internal/pokeapi"
	"github.com/ahgr3y/pokedex-repl-cli/internal/pokecache"
)

func main() {

	cfg := &config{
		pokeCache:     pokecache.NewPokeCache(),
		pokeapiClient: pokeapi.GetClient(time.Second*10, time.Minute*5),
	}

	startRepl(cfg)

}
