package collection

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type Bucket struct {
	Timestamp      int64
	DocumentKeySet *treeset.Set
}


func NewBucket(timestamp int64) (*Bucket) {
	return &Bucket{
		Timestamp:      timestamp,
		DocumentKeySet: treeset.NewWithStringComparator(),
	}
}


func (b *Bucket) Add(key string) {
	b.DocumentKeySet.Add(key)
}


func (b * Bucket) Remove(key string) {
	b.DocumentKeySet.Remove(key)
}


func (b *Bucket) Values() []interface{} {
	return b.DocumentKeySet.Values()
}


func (b *Bucket) Clear() {
	b.DocumentKeySet.Clear()
}


func (b *Bucket) ForEach(fnc func(int, interface{})) {
	b.DocumentKeySet.Each(fnc)
}


func (b *Bucket) Empty() (bool) {
	return b.DocumentKeySet.Empty()
}
