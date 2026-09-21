/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package util

import (
	"container/list"
	"sync"
)

// CacheMap 固定大小的 LRU 映射缓存
type CacheMap struct {
	maxSize int
	mu      sync.RWMutex
	lru     *list.List // map无序，需要list保存value
	items   map[interface{}]*list.Element
	/** 在get时检查元素是否失效，如果失效丢弃该元素 */
	needRemove func(interface{}) bool
	/**
     * 缓存对象被淘汰前的工作（如释放资源等），构造完后可以设置元素的淘汰按进入的先后次序淘汰最老元素，
     * 若淘汰时需要做一些额外操作（比如：清理资源）可以覆盖此方法
     */
	beforeRemove func(interface{})
}

// entry 存储在链表中的元素
type entry struct {
	key   interface{}
	value interface{}
}

// NewCacheMap 创建缓存实例
func NewCacheMap(maxSize int, needRemove func(interface{}) bool, beforeRemove func(interface{})) *CacheMap {
	return &CacheMap{
		maxSize:      maxSize,
		lru:          list.New(),
		items:        make(map[interface{}]*list.Element, maxSize),
		needRemove:   needRemove,
		beforeRemove: beforeRemove,
	}
}

// Get 从缓存中获取值
// 返回 (value, true) 如果命中
// 返回 (nil, false) 如果未命中
func (c *CacheMap) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return nil, false
	}
	e := elem.Value.(*entry)

	c.removeElement(elem)
	if c.needRemove(e.value) {
		c.beforeRemove(e.value)
		return nil, false
	}

	return e.value, true
}

// Put 将键值对放入缓存
func (c *CacheMap) Put(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 如果缓存已满，删除最老的元素
	if c.lru.Len() >= c.maxSize {
		elem := c.lru.Front()
		if elem != nil {
			c.beforeRemove(elem.Value.(*entry).value)
			c.removeElement(elem)
		}
	}

	// 插入新元素到链表尾部
	newEntry := &entry{key: key, value: value}
	elem := c.lru.PushBack(newEntry)
	c.items[key] = elem
}

// removeElement 从链表和 map 中移除元素
func (c *CacheMap) removeElement(elem *list.Element) {
	e := elem.Value.(*entry)
	delete(c.items, e.key)
	c.lru.Remove(elem)
}

// Size 返回当前缓存大小
func (c *CacheMap) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lru.Len()
}

// Clear 清空缓存
func (c *CacheMap) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lru.Init()
	c.items = make(map[interface{}]*list.Element, c.maxSize)
}

// Keys 返回所有 key（按从旧到新顺序）
func (c *CacheMap) Keys() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]interface{}, 0, c.lru.Len())
	for elem := c.lru.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(*entry)
		keys = append(keys, e.key)
	}
	return keys
}

// Values 返回所有 value（按从旧到新顺序）
func (c *CacheMap) Values() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]interface{}, 0, c.lru.Len())
	for elem := c.lru.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(*entry)
		values = append(values, e.value)
	}
	return values
}
