package gocache

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, data interface{}) error
	Remove(key string) error
	Flush()
}
