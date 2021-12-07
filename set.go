package gocache

func (cache *myCache) Set(key string, data interface{}) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	value := newValue(key, data)
	cache.items[key] = value

	return nil
}
