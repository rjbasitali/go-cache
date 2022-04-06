package gocache

import "time"

func (cache *myCache) deleteExpired() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	now := time.Now().UnixNano()
	for k, v := range cache.items {
		if cache.decisionFunc != nil && !cache.decisionFunc(v.data) {
			continue
		}
		if v.expiration > 0 && now > v.expiration {
			if cache.onEviction != nil {
				cache.onEviction(k, v.data)
			}
			delete(cache.items, k)
		}
	}
}
