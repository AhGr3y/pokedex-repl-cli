package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {

	internal := time.Second * 5
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {

		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(internal)
			cache.Add(c.key, c.val)
			value, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key: %v", c.key)
				return
			}
			if string(value) != string(c.val) {
				t.Errorf("Expected to find value: %v", string(value))
				return
			}

		})
	}

}

func TestReapLoop(t *testing.T) {

	const baseTime = time.Millisecond * 5
	const waitTime = baseTime + time.Millisecond*5
	cache := NewCache(baseTime)
	cache.Add("http://example.com", []byte("testdata"))
	_, ok := cache.Get("http://example.com")
	if !ok {
		t.Errorf("expecting to find a key")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("http://example.com")
	if ok {
		t.Errorf("expecting not to find a key")
	}

}
