package main

import (
	"fmt"
	"sync"
)

type Map[K comparable, V any] struct {
	m   map[K]V
	mtx sync.RWMutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m : make(map[K]V)}
}

func (m *Map[K, V]) Set(key K, value V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.m[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	val, ok := m.m[key]
	return val, ok
}

func (m *Map[K, V]) Delete(key K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.m, key)
}

func main() {
	safeMap := NewMap[string, int]()
	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", n), n*10)
		}(i)
	}

	wg.Wait()
	if val, ok := safeMap.Get("key9"); ok {
    	fmt.Println(val)
	} else {
    	fmt.Println("Ключ не найден")
	}	
}