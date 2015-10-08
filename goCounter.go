package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("hi")
}

type concurrentMap struct {
	m map[string]int
	*sync.RWMutex
}

func new() *concurrentMap {
	return &concurrentMap{make(map[string]int), &sync.RWMutex{}}
}

func (concurrentMap *concurrentMap) set(name string, count int) {
	concurrentMap.Lock()
	concurrentMap.m[name] = count
	concurrentMap.Unlock()
	return
}

func (concurrentMap *concurrentMap) get(name string) (count int) {
	concurrentMap.RLock()
	count = concurrentMap.m[name]
	concurrentMap.RUnlock()
	return count
}

func (concurrentMap *concurrentMap) increment(name string) (count int) {
	concurrentMap.Lock()
	concurrentMap.m[name]++
	count = concurrentMap.m[name]
	concurrentMap.Unlock()
	return count
}

func (concurrentMap *concurrentMap) getMap() map[string]int {
	return concurrentMap.m
}
