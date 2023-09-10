package hw04lrucache

import (
	"encoding/json"
	"fmt"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
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
	cashedItem, isExists := cache.items[key]
	newCacheItem := cacheItem{
		key:   key,
		value: value,
	}
	if isExists {
		cashedItem.Value = newCacheItem
		cache.queue.MoveToFront(cashedItem)
		debugMap(cache.items)
		return true
	}
	cache.queue.PushFront(newCacheItem)
	cache.items[key] = cache.queue.Front()
	if cache.queue.Len() > cache.capacity {
		lastUsedRecentItem := cache.queue.Back()
		cache.queue.Remove(lastUsedRecentItem)
		delete(cache.items, lastUsedRecentItem.Value.(cacheItem).key)
	}
	debugMap(cache.items)
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	item, isExists := cache.items[key]
	if !isExists {
		return nil, false
	}
	cache.queue.MoveToFront(item)
	return item.Value.(cacheItem).value, true
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func debugMap(collection map[Key]*ListItem) {
	bs, _ := json.Marshal(collection)
	fmt.Println(string(bs))
}
