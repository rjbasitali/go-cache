package gocache

import "time"

type value struct {
	key        string
	data       interface{}
	expiration int64
}

func newValue(key string, data interface{}, expiration int64) *value {
	return &value{
		key:        key,
		data:       data,
		expiration: time.Now().UnixNano() + expiration,
	}
}
