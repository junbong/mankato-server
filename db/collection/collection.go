package collection

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/Junbong/mankato-server/db/data"
)

type Collection struct {
	name    string
	data    hashmap.Map
	expire  uint16
}


func New(name string, expire uint16) (*Collection) {
	return &Collection{name:name, data:*hashmap.New(), expire:expire}
}


func (c Collection) String() string {
	return fmt.Sprintf("Collection{ name:%s, size:%d, expire:%d }", c.name, c.data.Size(), c.expire)
}


func (c Collection) Name() string {
	return c.name
}


func (c Collection) Put(key string, value string, expire uint16) {
	// TODO: make this function thread-safe

	//c.data.Put(key, value)
	d := data.New(key, value, expire)
	c.data.Put(key, d)
}


func (c Collection) Get(key string) (interface{}, bool) {
	// TODO: make this function thread-safe
	
	return c.data.Get(key)
}


func (c Collection) Remove(key string) bool {
	// TODO: make this function thread-safe
	
	before := c.data.Size()
	c.data.Remove(key)
	after := c.data.Size()
	return before != after
}


func (c Collection) Size() int {
	return c.data.Size()
}


func (c Collection) Clear() {
	// TODO: make this function thread-safe
	
	c.data.Clear()
}


func (c Collection) ContainsKey(key string) (contains bool) {
	// TODO: make this function thread-safe
	
	_, contains = c.data.Get(key)
	return
}
