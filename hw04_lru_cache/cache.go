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

func (lc *lruCache) Set(key Key, value interface{}) bool {
	item, heat := lc.items[key]
	if heat {
		item.Value = value
		lc.queue.MoveToFront(item)
	} else {
		if lc.queue.Len() == lc.capacity {
			for k, v := range lc.items {
				if v == lc.queue.Back() {
					delete(lc.items, k)
				}
			}
			lc.queue.Remove(lc.queue.Back())
		}
		lc.queue.PushFront(value)
	}
	lc.items[key] = lc.queue.Front()
	return heat
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	item, heat := lc.items[key]
	if heat {
		lc.queue.MoveToFront(item)
		lc.items[key] = lc.queue.Front()
		return lc.queue.Front().Value, heat
	}
	return nil, heat
}

func (lc *lruCache) Clear() {
	lc.items = make(map[Key]*ListItem, lc.capacity)
}
