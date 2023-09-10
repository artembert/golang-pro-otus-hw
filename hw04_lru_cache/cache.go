package hw04lrucache

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

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (list *lruCache) Set(key Key, value interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (list *lruCache) Get(key Key) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}

func (list *lruCache) Clear() {
	//TODO implement me
	panic("implement me")
}
