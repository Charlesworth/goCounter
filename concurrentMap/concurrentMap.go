package concurrentMap

import "sync"

type Map struct {
	m map[string]int
	*sync.RWMutex
}

func New() *Map {
	return &Map{make(map[string]int), &sync.RWMutex{}}
}

func (concurrentMap *Map) Set(name string, count int) {
	concurrentMap.Lock()
	concurrentMap.m[name] = count
	concurrentMap.Unlock()
	return
}

func (concurrentMap *Map) Get(name string) (count int) {
	concurrentMap.RLock()
	count = concurrentMap.m[name]
	concurrentMap.RUnlock()
	return count
}

func (concurrentMap *Map) Increment(name string) (count int) {
	concurrentMap.Lock()
	concurrentMap.m[name]++
	count = concurrentMap.m[name]
	concurrentMap.Unlock()
	return count
}

func (concurrentMap *Map) GetMap() map[string]int {
	concurrentMap.RLock()
	defer concurrentMap.RUnlock()
	return concurrentMap.m
}
