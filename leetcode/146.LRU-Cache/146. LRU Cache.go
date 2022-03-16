package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	mapping  map[int]*list.Element
	l        *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{mapping: map[int]*list.Element{}, l: list.New(), capacity: capacity}
	return cache
}

func (c *LRUCache) Get(key int) int {
	element, found := c.mapping[key]
	if !found {
		return -1
	}
	c.l.MoveToFront(element)
	return element.Value.([2]int)[1]
}

func (c *LRUCache) Put(key int, value int) {
	if element, found := c.mapping[key]; found {
		c.mapping[key].Value = [2]int{key, value}
		c.l.MoveToFront(element)
		return
	}
	c.l.PushFront([2]int{key, value})
	newElement := c.l.Front()
	c.mapping[key] = newElement

	for len(c.mapping) > c.capacity {
		backElement := c.l.Back()
		keyValue := backElement.Value.([2]int)
		backKey := keyValue[0]
		delete(c.mapping, backKey)
		c.l.Remove(backElement)
	}
}

func (c *LRUCache) print() {
	for key, element := range c.mapping {
		fmt.Printf("key: %d, value: %d ", key, element.Value.([2]int)[1])
	}
	println()
}

func main() {
	cache := Constructor(4)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Put(5, 5)
	cache.print()
}
