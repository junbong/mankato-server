package database

import (
	"time"
	"github.com/emirpasic/gods/maps/hashmap"
	"fmt"
	"github.com/Junbong/mankato-server/db/collection"
	"log"
)

type Database struct {
	createdAt   time.Time
	collections hashmap.Map
}


func New() (*Database) {
	return &Database{createdAt:time.Now(), collections:*hashmap.New()}
}


func (d Database) String() string {
	return fmt.Sprintf("Database{ createdAt:%s, size:%d }", d.createdAt, d.collections.Size())
}


func (d Database) Put(nameOfCollection string, col interface{}) {
	// TODO: make this function thread-safe
	
	d.collections.Put(nameOfCollection, col)
}


func (d Database) Get(nameOfCollection string) (col interface{}, exists bool) {
	// TODO: make this function thread-safe
	
	col, exists = d.collections.Get(nameOfCollection)
	
	if exists {
		col = col.(*collection.Collection)
	}
	
	return
}


func (d Database) GetOrCreateCollection(nameOfCollection string, createIfNotExists bool) (col interface{}) {
	// TODO: make this function thread-safe
	
	var exists bool
	col, exists = d.collections.Get(nameOfCollection)
	
	if !exists && createIfNotExists {
		col = collection.New(nameOfCollection)
		d.Put(nameOfCollection, col)
		log.Println(
			"New collection added [", nameOfCollection,
			"] Total number of collections =", d.collections.Size())
	}
	
	return
}


func (d Database) Remove(nameOfCollection string) bool {
	// TODO: make this function thread-safe

	before := d.collections.Size()
	d.collections.Remove(nameOfCollection)
	after := d.collections.Size()
	return before != after
}


//func (c Collection) Clear() {
//	// TODO: make this function thread-safe
//
//	c.data.Clear()
//}


//func (c Collection) ContainsKey(key string) (contains bool) {
//	// TODO: make this function thread-safe
//
//	_, contains = c.data.Get(key)
//	return
//}
