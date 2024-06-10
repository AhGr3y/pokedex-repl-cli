package main

import (
	"time"

	"github.com/ahgr3y/pokedex-repl-cli/internal/pokeapi"
)

func main() {

	cfg := &config{
		pokeapiClient: pokeapi.GetClient(time.Second*10, time.Minute*5),
	}

	startRepl(cfg)

}
