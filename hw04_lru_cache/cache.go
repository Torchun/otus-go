package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	//Cache // Remove me after realization.

	capacity int
	queue    List // type from list.go with described methods
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	// set means key goes to front of cache regardless of
	// check if key in cache
	_, ok := l.items[key]
	if !ok {
		// check & remove exceeded cache capacity item
		if l.queue.Len() >= l.capacity {
			delete(l.items, l.queue.Back().Key) // remove exceeded key from map
			l.queue.Remove(l.queue.Back())      // remove last item from cache
		}
		// capacity is not exceeded, simply add key to map and item to front of cache
		l.items[key] = l.queue.PushFront(value)
		l.items[key].Key = key
		return false // was not found in cache
	}
	// key found in cache, need to update value
	l.queue.Remove(l.items[key])
	delete(l.items, key)

	l.items[key] = l.queue.PushFront(value)
	l.items[key].Key = key
	return true // it was found in cache
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(item)
		l.items[key].Key = key
		l.queue.Front().Key = key

		return item.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem)
	l.queue = nil
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
