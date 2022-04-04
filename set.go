package gocache

import "time"

func (cache *myCache) Set(key string, data interface{}) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	value := newValue(key, data, time.Now().UnixNano()+cache.expireAfter)
	cache.items[key] = value

	return nil
}
