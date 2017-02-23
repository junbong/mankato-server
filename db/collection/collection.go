package collection

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/Junbong/mankato-server/db/documents"
)

type Collection struct {
	Name      string
	Documents hashmap.Map
}


func New(name string) (*Collection) {
	return &Collection{Name:name, Documents:*hashmap.New()}
}


func (c Collection) String() string {
	return fmt.Sprintf("Collection{ name:%s, size:%d }", c.Name, c.Documents.Size())
}


func (c Collection) Put(key string, value string, expire uint16) {
	// TODO: make this function thread-safe

	d := data.New(key, value, expire)
	c.Documents.Put(key, d)
}


func (c Collection) Get(key string) (interface{}, bool) {
	// TODO: make this function thread-safe
	
	return c.Documents.Get(key)
}


func (c Collection) Remove(key string) bool {
	// TODO: make this function thread-safe
	
	before := c.Documents.Size()
	c.Documents.Remove(key)
	after := c.Documents.Size()
	return before != after
}


func (c Collection) Size() int {
	return c.Documents.Size()
}


func (c Collection) Clear() {
	// TODO: make this function thread-safe
	
	c.Documents.Clear()
}


func (c Collection) ContainsKey(key string) (contains bool) {
	// TODO: make this function thread-safe
	
	_, contains = c.Documents.Get(key)
	return
}
