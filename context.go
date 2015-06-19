package stack

import (
	"fmt"
	"sync"
)

type Context struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

func NewContext() *Context {
	m := make(map[string]interface{})
	return &Context{m: m}
}

func (c *Context) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val := c.m[key]
	if val == nil {
		return nil, fmt.Errorf("stack.Context: key '%s' does not exist", key)
	}
	return val, nil
}

func (c *Context) Put(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = val
}

func (c *Context) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}