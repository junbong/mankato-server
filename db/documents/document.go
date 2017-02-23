package data

import "fmt"

type Document struct {
	Key       string
	Value     string
	CreatedAt uint16
	ExpiresAt uint16
}


func New(key string, value string, expire uint16) (*Document) {
	// TODO: 'expire' need to be changed to 'ttl'
	return &Document{Key:key, Value:value, ExpiresAt:expire}
}


func (d Document) String() string {
	return fmt.Sprintf("Document{ key:%s, value:%s, expired_at:%d }", d.Key, d.Value, d.ExpiresAt)
}
