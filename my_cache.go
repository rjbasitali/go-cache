package gocache

import (
	"sync"
)

type myCache struct {
	mutex sync.Mutex
	items map[string]*value
}

func NewCache() Cache {
	return &myCache{
		items: make(map[string]*value),
	}
}
