package gocache

import "time"

func (cache *myCache) deleteExpired() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	now := time.Now().UnixNano()
	for k, v := range cache.items {
		if cache.decisionFunc != nil && !cache.decisionFunc(v) {
			continue
		}
		if v.expiration > 0 && now > v.expiration {
			delete(cache.items, k)
		}
	}
}
