package document

import (
	"fmt"
	"time"
)

type Document struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt int64     `json:"_created_at"`
	ExpiresAt int64     `json:"_expires_at"`
}


func New(key string, value string, expAfterSec uint) (*Document) {
	now := time.Now()
	exp := now.Add(time.Second * time.Duration(expAfterSec))
	return &Document{Key:key, Value:value, CreatedAt:now.Unix(), ExpiresAt:exp.Unix()}
}


func (d Document) String() string {
	return fmt.Sprintf("Document{ key:%s, value:%s, expired_at:%d }", d.Key, d.Value, d.ExpiresAt)
}
