package gocache

type value struct {
	key  string
	data interface{}
}

func newValue(key string, data interface{}) *value {
	return &value{
		key:  key,
		data: data,
	}
}
