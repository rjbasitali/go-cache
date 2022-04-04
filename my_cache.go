package gocache

import (
	"runtime"
	"sync"
	"time"
)

type myCache struct {
	mutex       sync.Mutex
	items       map[string]*value
	expireAfter int64
}

func NewCache() Cache {
	return &myCache{
		items: make(map[string]*value),
	}
}

func NewCacheWithSweeper(interval, expireAfter time.Duration) Cache {
	c := &myCache{
		items:       make(map[string]*value),
		expireAfter: expireAfter.Nanoseconds(),
	}

	if interval > 0 && expireAfter > 0 {
		s := newSweeper(c, interval)
		runtime.SetFinalizer(c, s.stopSweeper)
		go s.run(c)
	}

	return c
}
