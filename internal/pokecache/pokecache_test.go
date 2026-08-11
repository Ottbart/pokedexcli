package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	t.Logf("Starting TestAddGet with interval: %v", interval)
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
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Error("expected to find key")
				return
			}
			t.Logf("Key found in cache")

			if string(val) != string(c.val) {
				t.Error("expected to find value")
				return
			}
			t.Logf("Value matches expected value")
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	t.Logf("Starting TestReapLoop with baseTime: %v and waitTime: %v", baseTime, waitTime)
	cache := NewCache(baseTime)
	t.Logf("Adding key to cache")
	cache.Add("https://example.com", []byte("testdata"))
	t.Logf("Testing cache.Get for key")
	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Error("expected to find key")
		return
	}
	t.Logf("Key found in cache")

	time.Sleep(waitTime)
	t.Logf("Testing cache.Get for key after waiting for reapLoop to run")
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Error("expected to not find key")
		return
	}
	t.Logf("Key not found in cache as expected")

}
