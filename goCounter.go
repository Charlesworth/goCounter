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

func (c *concurrentMap) set(name string, count int) {
	c.Lock()
	c.m[name] = count
	c.Unlock()
	return
}

func (c *concurrentMap) get(name string) (count int) {
	c.RLock()
	count = c.m[name]
	c.RUnlock()
	return count
}
