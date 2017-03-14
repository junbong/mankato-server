package document

import (
	"fmt"
	"time"
	"github.com/Junbong/mankato-server/utils"
	"regexp"
	"log"
)

const (
	DEFAULT_CONTENT_TYPE string = "text/plain"
	CONTENT_TYPE_REGEX string   = "[A-z0-9]+\\/[A-z0-9]+"
	NEVER_EXPIRES string        = "never"
)

type Document struct {
	Key         string          `json:"key"`
	Value       []byte          `json:"value"`
	ContentType string          `json:"content_type"`
	Expired     bool            `json:"_expired"`
	CreatedAt   interface{}     `json:"_created_at"`
	ExpiresAt   interface{}     `json:"_expires_at"`
}


func New(key string, value []byte, contentType string, expAfterSec int) (*Document) {
	now := time.Now()
	
	// Content type
	match, err := regexp.MatchString(CONTENT_TYPE_REGEX, contentType)
	if !match || err != nil {
		log.Printf("Unresolved content type '%s' so use default content type %s",
			contentType, DEFAULT_CONTENT_TYPE)
		contentType = DEFAULT_CONTENT_TYPE
	}
	
	// Create new document
	d := Document{Key:key, ContentType:contentType, CreatedAt:now.Unix()}
	
	// Set value
	d.updateValueAndExpires(value, expAfterSec, now)
	
	return &d
}


func (d *Document) String() string {
	var exAt string
	
	if d.ExpiresAt != nil {
		exAt = time.Unix(d.ExpiresAt.(int64), 0).String()
	} else {
		exAt = NEVER_EXPIRES
	}
	
	return fmt.Sprintf("Document{ key: %s, value: %s, content_type: %s, created_at: %s expires_at: %s }",
		d.Key, d.Value, d.ContentType, time.Unix(d.CreatedAt.(int64), 0).String(), exAt)
}


func (d *Document) Update(value []byte, expAfterSec int) (*Document) {
	return d.updateValueAndExpires(value, expAfterSec, time.Now())
}


func (d *Document) updateValueAndExpires(
		value []byte,
		expAfterSec int,
		expStdTime time.Time) (*Document) {
	// Update value
	if utils.IsNotNilOrEmpty(value) {
		d.Value = make([]byte, len(value))
		copy(d.Value, value)
	} else {
		// Empty slice
		d.Value = make([]byte, 0)
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


/*
Make this document expired.
 */
func (d *Document) Expire() (*Document) {
	d.Expired = true
	
	fmt.Printf("Document Expired: %s \n", d.Key)
	
	return d
}
