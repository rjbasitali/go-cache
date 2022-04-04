package gocache

func (cache *myCache) Flush() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.items = make(map[string]*value)
}
