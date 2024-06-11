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
