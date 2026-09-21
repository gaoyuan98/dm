/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package util

import (
	"container/list"
	"sync"
)

// CacheQueue 固定大小的队列缓存
type CacheQueue struct {
	maxSize      int
	mu           sync.Mutex
	list         *list.List
	enableLRU    bool
	beforeRemove func(interface{})
}

// NewLRU 创建 LRU 缓存
func NewCacheQueue(maxSize int, enableLRU bool, beforeRemove func(interface{})) *CacheQueue {
	return &CacheQueue{
		maxSize:      maxSize,
		list:         list.New(),
		enableLRU:    enableLRU,
		beforeRemove: beforeRemove,
	}
}

// Get 获取值
func (c *CacheQueue) Get() (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem := c.list.Front()
	if elem == nil {
		return nil, false
	}
	c.list.Remove(elem)
	return elem.Value, true
}

// Put 放入缓存
// 如果容量已满，则删除队首（最老）元素
func (c *CacheQueue) Put(value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 如果已存在，则直接返回
	for elem := c.list.Front(); elem != nil; elem = elem.Next() {
		if elem.Value == value {
			return true
		}
	}
	if c.list.Len() < c.maxSize {
		// 新元素插入队尾
		c.list.PushBack(value)
		return true
	} else if c.enableLRU {
		elem := c.list.Front()
		c.beforeRemove(elem.Value)
		c.list.Remove(elem)
		c.list.PushBack(value)
		return true
	} else {
		c.beforeRemove(value)
		return false
	}
}

// Contains 检查 value 是否存在
func (c *CacheQueue) Contains(value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	for elem := c.list.Front(); elem != nil; elem = elem.Next() {
		if elem.Value == value {
			return true
		}
	}
	return false
}

// Size 返回当前大小
func (c *CacheQueue) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.list.Len()
}

// Clear 清空缓存
func (c *CacheQueue) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.list.Init()
}

// Values 返回所有 value（从旧到新）
func (c *CacheQueue) Values() []interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	values := make([]interface{}, 0, c.list.Len())
	for elem := c.list.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value)
	}
	return values
}