package mathbeta

import (
	_ "log"
	_ "fmt"
	"sync"
)

type Xcache struct {
	Cache map[string]interface{}
	sync.RWMutex
}

func (c *Xcache) Put(key string, value interface{}) bool {
	c.Cache[key]=value
	return true
}

func (c *Xcache) Get(key string) interface{} {
	return c.Cache[key]
}

func (c *Xcache) Delete(key string) bool {
	delete(c.Cache, key)
	_, exists:=c.Cache[key]
	return exists;
}

func (c *Xcache) Cas(key string, oldValue, newValue interface{}) bool {
	c.Lock()

	f:=false
	currValue:=c.Get(key)
	if currValue==oldValue {
		c.Cache[key]=newValue
		f=true
	}

	c.Unlock()
	return f
}

func (c *Xcache) Getsize() int {
	return len(c.Cache)
}