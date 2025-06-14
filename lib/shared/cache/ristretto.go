package cache

import "github.com/dgraph-io/ristretto/v2"

type RistrettoCache struct {
	Cache *ristretto.Cache[string, any]
}

var GLOBAL *RistrettoCache

func GetRistrettoCache() *RistrettoCache {
	if GLOBAL != nil {
		return GLOBAL
	}

	cache, err := ristretto.NewCache(&ristretto.Config[string, any]{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})

	if err != nil {
		panic(err)
	}

	GLOBAL = &RistrettoCache{
		Cache: cache,
	}

	return GLOBAL
}
