package cache

import (
	"time"
)

type Item struct {
	Value    string
	deadline *time.Time
}

type Cache struct {
	items map[string]Item
}

func NewCache() Cache {
	return Cache{
		items: map[string]Item{},
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	_, ok := cache.items[key]
	if !ok {
		return "", false
	}

	now := time.Now()
	if cache.items[key].deadline != nil && cache.items[key].deadline.Before(now) {
		return "", false
	}

	return cache.items[key].Value, true
}

func (cache *Cache) Put(key, value string) {
	cache.items[key] = Item{Value: value, deadline: nil}
}

func (cache *Cache) Keys() []string {
	var keys []string
	now := time.Now()

	for k, v := range cache.items {
		if v.deadline != nil && v.deadline.Before(now) {
			continue
		}

		keys = append(keys, k)
	}

	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.items[key] = Item{value, &deadline}
}
