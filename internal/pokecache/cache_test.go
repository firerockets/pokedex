package pokecache

import (
	"testing"
	"time"
)

const duration time.Duration = 2 * time.Microsecond

func TestAddGet(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "test.com",
			value: []byte("test"),
		},
		{
			key:   "abc.com",
			value: []byte("abc"),
		},
	}

	cache := NewCache(duration)

	for _, c := range cases {
		cache.Add(c.key, c.value)

		val, exists := cache.Get(c.key)

		if exists == false {
			t.Errorf("Expected to have some value to added key %s", c.key)
			t.Fail()
		}

		if val == nil {
			t.Errorf("Expected values to match. %s not equal to %s.", val, c.value)
			t.Fail()
		}
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(duration)

	key := "test.com"
	cache.Add(key, []byte("test"))

	_, ok := cache.Get(key)

	if !ok {
		t.Error("Expected to find some value")
		t.Fail()
	}

	time.Sleep(duration)

	_, ok = cache.Get(key)

	if ok {
		t.Error("Expected to NOT find any value")
		t.Fail()
	}
}
