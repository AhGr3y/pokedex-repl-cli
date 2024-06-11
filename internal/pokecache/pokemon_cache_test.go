package pokecache

import (
	"fmt"
	"testing"

	"github.com/ahgr3y/pokedex-repl-cli/internal/poketype"
)

func TestPokeCacheAddGet(t *testing.T) {
	cases := []struct {
		key string
		val poketype.Pokemon
	}{
		{
			key: "pikachu",
			val: poketype.Pokemon{
				Name: "pikachu",
			},
		},
		{
			key: "pidgey",
			val: poketype.Pokemon{
				Name: "pidgey",
			},
		},
	}

	for i, c := range cases {

		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewPokeCache()
			cache.Add(c.key, c.val)
			value, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key: %v", c.key)
				return
			}
			if value.Name != c.val.Name {
				t.Errorf("%v != %v, Expected to find: %v", value, c.val, c.val)
			}
		})
	}
}

func TestLen(t *testing.T) {
	cache := NewPokeCache()
	if cache.Len() != 0 {
		t.Errorf("Expected length to be 0")
	}
	cache.Add("pikachu", poketype.Pokemon{})
	if cache.Len() != 1 {
		t.Errorf("Expected length to be 1")
	}
}

func TestGetPokemonList(t *testing.T) {
	cache := NewPokeCache()
	cache.Add("pikachu", poketype.Pokemon{})
	cache.Add("pidgey", poketype.Pokemon{})
	list := cache.GetPokemonList()
	if list[0] != "pikachu" {
		t.Errorf("Expected to find pikachu")
	}
	if list[1] != "pidgey" {
		t.Errorf("Expected to find pidgey")
	}
}
