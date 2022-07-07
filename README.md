# go-cache

Simple thread safe key/value cache for Go.

## Usage

You can initialize the cache using:

```go
    cache = gocache.NewCache()
```

And perform the operations using:

```go
   // GET a value using key, which returns (value, error)
   cache.Get(key)
   
   // SET a value using key
   cache.Set(key, value)
   
   // REMOVE a key/value pair using key
   cache.Remove(key)
   
   // And FLUSH, which removes all items from the cache
   cache.Flush()
```

## Cache Expiration

You can also add an **expiration** to the cache, expired cache will be auto deleted by the `sweeper`.

To initialize a cache using `sweeper` we can use:

```go
    cache := NewCacheWithSweeper(10 * time.Minute, 10 * time.Minute, decisionFunc, onEviction)
```

Here, first two parameters are `interval` and `expireAfter`, sweeper will run in the background after specified `interval` to check for expired cache to delete.

The `decisionFunc` and `onEviction` are **optional** functions with the following signature:
```go
    decisionFunc func(v interface{}) bool
    
    onEviction func(k string, v interface{})
```

We can use `decisionFunc` as a callback, it will be called by the cache before deleting a key/value, it takes in the value to delete and returns `true` to delete and `false` for not deleting a key/value pair even after expired.

The `onEviction` is a callback which will be called every time a cache is deleted with its key/value as parameters.
