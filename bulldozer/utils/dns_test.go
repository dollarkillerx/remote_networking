package utils

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func TestDns(t *testing.T) {
	dns, err := SearchDns("www.baidu.com", "8.8.8.8:53")
	if err != nil {
		return
	}

	marshal, err := json.Marshal(dns)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}

func TestCache(t *testing.T) {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", time.Hour)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		f, ok := foo.(string)
		if ok {
			fmt.Println(f)
		}
	}

}
