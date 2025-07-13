//go:build !solution

package lrucache

import "container/list"

type Item struct {
	key   int
	value int
}

type LRU struct {
	cap   int
	m     map[int]*list.Element
	items *list.List
}

func (l *LRU) Clear() {
	for k, v := range l.m {
		delete(l.m, k)
		l.items.Remove(v)
	}
}

func (l *LRU) Get(key int) (int, bool) {
	if el, ok := l.m[key]; ok {
		l.items.MoveToFront(el)
		return el.Value.(*Item).value, ok
	}

	return 0, false
}

func (l *LRU) Set(key, value int) {
	if el, ok := l.m[key]; len(l.m) >= l.cap && !ok {
		if l.items.Len() == 0 {
			return
		}
		lastItem := l.items.Back().Value.(*Item)
		delete(l.m, lastItem.key)
		l.items.Remove(l.items.Back())
	} else if ok {
		l.items.Remove(el)
		delete(l.m, key)
	}
	item := &Item{key, value}
	el := l.items.PushFront(item)
	l.m[key] = el
}

func (l *LRU) Range(f func(key, value int) bool) {
	for el := l.items.Back(); el != nil; el = el.Prev() {
		item := el.Value.(*Item)
		shouldContinue := f(item.key, item.value)
		if !shouldContinue {
			return
		}
	}
}

func New(cap int) Cache {
	return &LRU{
		cap:   cap,
		m:     make(map[int]*list.Element, cap),
		items: list.New(),
	}
}
