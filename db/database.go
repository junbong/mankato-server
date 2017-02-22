package db

import (
	"time"
	"log"
	"github.com/emirpasic/gods/maps/hashmap"
)

type Database struct {
	createdAt   time.Time
	collections hashmap.Map
}

type Collection struct {
	name    string
	data    hashmap.Map
	expire  uint16
}

type Data struct {
	key     string
	value   string
	expire  uint16
}


func New() (db *Database) {
	db = &Database{createdAt: time.Now(), collections: *hashmap.New()}
	log.Println("Database initialized")
	return
}


func (db *Database) Put(nameOfCollection, key, value string) {
	// TODO: need raise error when nameOfCollection or key is nil
	
	data := Data{key:key, value:value, expire:0}
	collection, _ := db.GetCollection(nameOfCollection, true)
	collection.data.Put(key, data)
}


func (db *Database) GetCollection(nameOfCollection string, createIfNotExists bool) (*Collection, bool) {
	// TODO: need raise error when nameOfCollection is nil
	
	var collection interface{}
	var exists bool
	
	// If collection is not exists, create new one with specified name
	if collection, exists = db.collections.Get(nameOfCollection);
			!exists && createIfNotExists {
		collection = createCollection(nameOfCollection, 0)
		db.collections.Put(nameOfCollection, collection)
		log.Println(
			"New collection added [", nameOfCollection,
			"] Total number of collections =", db.collections.Size())
	}
	
	if collection != nil {
		return collection.(*Collection), true
	} else {
		return nil, false
	}
}


func (db *Database) SizeOfCollection(collection *Collection) int {
	return collection.data.Size()
}


/* Private - create collection with given name and expiration time */
func createCollection(name string, expire uint16) *Collection {
	collection := Collection{name:name, data:*hashmap.New(), expire:expire}
	return &collection
}
