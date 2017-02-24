package document

import (
	"fmt"
	"time"
	"github.com/Junbong/mankato-server/utils"
)

type Document struct {
	Key       string        `json:"key"`
	Value     string        `json:"value"`
	CreatedAt interface{}   `json:"_created_at"`
	ExpiresAt interface{}   `json:"_expires_at"`
}


func New(key string, value string, expAfterSec int) (*Document) {
	now := time.Now()
	
	d := Document{Key:key, CreatedAt:now.Unix()}
	d.updateValueAndExpires(value, expAfterSec, now)
	
	return &d
}


func (d *Document) String() string {
	var exAt string
	
	if d.ExpiresAt != nil {
		exAt = time.Unix(d.ExpiresAt.(int64), 0).String()
	} else {
		exAt = "never"
	}
	
	return fmt.Sprintf("Document{ key:%s, value:%s, created_at:%s expires_at:%s }",
		d.Key, d.Value, time.Unix(d.CreatedAt.(int64), 0).String(), exAt)
}


func (d *Document) Update(value string, expAfterSec int) (*Document) {
	return d.updateValueAndExpires(value, expAfterSec, time.Now())
}


func (d *Document) updateValueAndExpires(
		value string,
		expAfterSec int,
		expStdTime time.Time) (*Document) {
	// Update value
	if utils.IsNotNilOrEmpty(value) {
		d.Value = value
	} else {
		d.Value = ""
	}
	
	// Update TTL
	if expAfterSec >= 0 {
		d.ExpiresAt = expStdTime.
			Add(time.Second * time.Duration(expAfterSec)).
			Unix()
	} else {
		d.ExpiresAt = nil
	}
	
	return d
}
