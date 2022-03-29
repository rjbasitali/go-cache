package gocache

func (cache *myCache) Get(key string) (interface{}, error) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	item, exists := cache.items[key]
	if !exists {
		return nil, ErrKeyNotFound
	}

	return item.data, nil
}
