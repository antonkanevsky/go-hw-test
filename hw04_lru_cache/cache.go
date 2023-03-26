package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.Lock()
	defer cache.Unlock()
	listItem, itemExists := cache.items[key]
	if itemExists {
		cacheItem := listItem.Value.(*cacheItem)
		cacheItem.value = value
		cache.queue.MoveToFront(listItem)
	} else {
		cache.items[key] = cache.queue.PushFront(&cacheItem{
			key:   key,
			value: value,
		})
	}

	if cache.queue.Len() > cache.capacity {
		lastListItem := cache.queue.Back()
		lastCacheItem := lastListItem.Value.(*cacheItem)
		delete(cache.items, lastCacheItem.key)
		cache.queue.Remove(lastListItem)
	}

	return itemExists
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.Lock()
	defer cache.Unlock()
	listItem, itemExists := cache.items[key]
	if itemExists {
		cache.queue.MoveToFront(listItem)
		listCacheItem := listItem.Value.(*cacheItem)

		return listCacheItem.value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}
