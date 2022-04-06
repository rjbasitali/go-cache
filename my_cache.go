package gocache

import (
	"runtime"
	"sync"
	"time"
)

type myCache struct {
	mutex        sync.Mutex
	items        map[string]*value
	expireAfter  int64
	decisionFunc func(v interface{}) bool
	onEviction   func(k string, v interface{})
}

func NewCache() Cache {
	return &myCache{
		items: make(map[string]*value),
	}
}

func NewCacheWithSweeper(interval, expireAfter time.Duration,
	decisionFunc func(v interface{}) bool, onEviction func(k string, v interface{})) Cache {
	c := &myCache{
		items:        make(map[string]*value),
		expireAfter:  expireAfter.Nanoseconds(),
		decisionFunc: decisionFunc,
		onEviction:   onEviction,
	}

	if interval > 0 && expireAfter > 0 {
		s := newSweeper(c, interval)
		runtime.SetFinalizer(c, s.stopSweeper)
		go s.run(c)
	}

	return c
}
