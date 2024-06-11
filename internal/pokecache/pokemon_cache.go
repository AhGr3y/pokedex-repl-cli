package pokecache

import (
	"sync"

	"github.com/ahgr3y/pokedex-repl-cli/internal/poketype"
)

// who will use pokemonCache struct?
// commandCatch -> main package
// catchPokemon funtion -> pokeapi package

type PokemonCache struct {
	pokeCache map[string]poketype.Pokemon
	mu        *sync.Mutex
}

func (c *PokemonCache) Add(key string, val poketype.Pokemon) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pokeCache[key] = val
}

func (c *PokemonCache) Get(key string) (poketype.Pokemon, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	pokemon, ok := c.pokeCache[key]
	return pokemon, ok
}

func (c *PokemonCache) Len() int {
	return len(c.pokeCache)
}

func (c *PokemonCache) GetPokemonList() []string {
	if c.Len() == 0 {
		return []string{}
	}

	list := []string{}
	for name := range c.pokeCache {
		list = append(list, name)
	}
	return list
}

func NewPokeCache() *PokemonCache {
	return &PokemonCache{
		pokeCache: make(map[string]poketype.Pokemon),
		mu:        &sync.Mutex{},
	}
}
