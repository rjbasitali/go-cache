package gocache

func (cache *myCache) Remove(key string) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	_, exists := cache.items[key]
	if !exists {
		return ErrKeyNotFound
	}

	delete(cache.items, key)

	return nil
}
