package db

type BasicKeyValueDataStructure interface {
	New() (*BasicKeyValueDataStructure)
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, exists bool)
	Remove(key interface{}) bool
	Size() int
	Clear()
	ContainsKey(key interface{}) (contains bool)
}
