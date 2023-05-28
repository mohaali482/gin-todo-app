package caches

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type CacheInterface interface {
	Set(key string, value interface{})
	Get(key string, value interface{}) error
	Delete(key string) error
	DeleteAllTodoCache()
}

func get_ring() *redis.Ring {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": "localhost:6379",
		},
	})

	return ring
}

func get_cache() *cache.Cache {
	ring := get_ring()
	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return mycache
}

func Set(key string, value interface{}) {
	my_cache := get_cache()
	ctx := context.TODO()

	if err := my_cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}
}

func Get(key string, value interface{}) error {
	my_cache := get_cache()
	ctx := context.TODO()
	if err := my_cache.Get(ctx, key, &value); err != nil {
		if err == cache.ErrCacheMiss {
			return cache.ErrCacheMiss
		}
		panic(err)
	}

	return nil
}

func Delete(key string) {
	my_cache := get_cache()
	ctx := context.TODO()
	if err := my_cache.Delete(ctx, key); err != nil {
		panic(err)
	}
}

func DeleteAllTodoCache() {
	my_cache := get_cache()
	ctx := context.TODO()
	iter := get_ring().Scan(ctx, 0, "todo:*", 0).Iterator()
	for iter.Next(ctx) {
		err := my_cache.Delete(ctx, iter.Val())
		if err != nil {
			panic(err)
		}
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

}
