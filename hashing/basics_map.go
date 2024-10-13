package hashing

import (
	"fmt"
	"sync"
)

// Go Map implementation https://github.com/golang/go/blob/master/src/runtime/map.go#L7-L54

var inBuildSyncMap = sync.Map{}

type Map struct {
	store map[string]string
	sync.Mutex
}

func NewMap() *Map {
	return &Map{
		store: make(map[string]string),
	}
}

func (m *Map) Set(key, value string) {
	m.Lock()
	defer m.Unlock()
	m.store[key] = value
}

func (m *Map) Get(key string) (string, bool) {
	m.Lock()
	defer m.Unlock()
	value, ok := m.store[key]
	return value, ok
}

func (m *Map) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.store, key)
}

func (m *Map) Clear() {
	m.Lock()
	defer m.Unlock()
	m.store = make(map[string]string)
}

func (m *Map) Len() int {
	m.Lock()
	defer m.Unlock()
	return len(m.store)
}

func InBuildSyncMapSet(key, value string) {
	// Store value in sync.Map
	inBuildSyncMap.Store(key, value)

	// Range over the sync.Map
	inBuildSyncMap.Range(func(key, value interface{}) bool {
		// do something with key and value
		return true
	})

	// Load value from sync.Map
	if v, ok := inBuildSyncMap.Load(key); ok {
		// do something with val
		fmt.Println(v)
	}

	// Delete value from sync.Map
	inBuildSyncMap.Delete(key)
}
