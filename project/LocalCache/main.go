package main

import (
	"fmt"
	"github.com/coocood/freecache"
)

func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	key := []byte("k1")
	value := []byte("v1")
	expire := 60
	cache.Set(key, value, expire)

	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("got -->", string(got))
	}
}
