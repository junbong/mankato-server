package collection

import (
	"github.com/Junbong/mankato-server/configs"
	"github.com/Junbong/mankato-server/db/documents"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"
	"github.com/Junbong/mankato-server/exphandlers"
	"github.com/Junbong/mankato-server/exphandlers/amqp"
	"fmt"
	"time"
	"sync"
	"log"
)

type Collection struct {
	Name                string
	Opened              bool
	ScanTtl             bool
	Documents           hashmap.Map
	TtlIndex            btree.Tree
	OpLock              sync.Mutex
	CreatedAt           int64
	ttlScanPeriodMillis int
	expirationHandler   []exphandler.ExpirationHandler
}

const (
	// It will replaced when value of configuration is not same with this value
	TTL_SCANNER_DELAY_MILLIS = 1000
)


func New(name string, config *configs.Config) (*Collection) {
	c := &Collection{
		Name:       name,
		ScanTtl:    true,
		Documents:  *hashmap.New(),
		TtlIndex:   *btree.NewWith(3, utils.Int64Comparator),
		CreatedAt:  time.Now().Unix(),
	}
	
	// Configs
	if config.Collection.TtlScanPeriodMillis != TTL_SCANNER_DELAY_MILLIS {
		c.ttlScanPeriodMillis = config.Collection.TtlScanPeriodMillis
	} else {
		c.ttlScanPeriodMillis = TTL_SCANNER_DELAY_MILLIS
	}
	
	// Expiration Handlers
	for _, hcfg := range config.ExpirationHandler {
		if hcfg.Enable {
			switch hcfg.Type {
			case "amqp":
				h := amqpexphandler.New(
					hcfg.Properties["uri"],
					hcfg.Properties["queue"],
				)
				if err := h.Open(); err == nil {
					c.expirationHandler = append(c.expirationHandler, h)
				} else {
					log.Fatalln("Expiration handler is not opened ", h)
				}
				
			default:
				panic(fmt.Sprintf("Invalid handler type '%s'", hcfg.Type))
			}
		}
	}
	
	return c
}


func (c *Collection) String() string {
	return fmt.Sprintf(
		"Collection{ name:%s, size:%d, open:%v, created_at:%s }",
		c.Name, c.Documents.Size(), c.Opened, time.Unix(c.CreatedAt, 0))
}


func (c *Collection) Open() (*Collection) {
	if c.ScanTtl {
		go func(col *Collection) {
			for col.Opened {
				time.Sleep(time.Millisecond * time.Duration(c.ttlScanPeriodMillis))
				c.scanTtlsAndRemove()
			}
		}(c)
	}
	
	c.Opened = true
	
	log.Printf("Collection [ %s ] opened", c.Name)
	return c
}


func (c *Collection) checkOpened() {
	if !c.Opened {
		panic(fmt.Sprintf(
			"Collection [ %s ] is not opened! Did you call Collection.Open()?",
			c.Name))
	}
}


func (c *Collection) scanTtlsAndRemove() {
	now := time.Now().Unix()
	
	for true {
		min := c.TtlIndex.LeftKey()
		
		if min != nil && min.(int64) <= now {
			// Remove all keys in Bucket
			bucket, _ := c.TtlIndex.Get(min);
			bucket.(*Bucket).ForEach(func(i int, k interface{}) {
				if d, exists := c.Documents.Get(k); exists {
					docT := d.(*document.Document)
					docT.Expire()
					
					// Remove Document
					c.Documents.Remove(docT.Key)
					
					// Handle expiration
					for _, handler := range c.expirationHandler {
						handler.HandleDocument(docT)
					}
				}
			})
			bucket.(*Bucket).Clear()
			
			// Remove Bucket in index
			c.TtlIndex.Remove(min)
			
		} else {
			break
		}
	}
}


func (c *Collection) Put(
		key string,
		value string,
		expAfterSec int) (*document.Document) {
	c.checkOpened()

	doc := document.New(key, value, expAfterSec)
	
	c.OpLock.Lock()
	c.Documents.Put(key, doc)
	
	// Add to index
	c.putOrUpdateTtlIndex(doc, nil)
	defer c.OpLock.Unlock()
	
	return doc
}


func (c *Collection) PutOrUpdate(
		key string,
		value string,
		expAfterSec int) (*document.Document) {
	c.checkOpened()
	
	// Already has key, update it
	if doc, exists := c.Get(key); exists {
		docT := doc.(*document.Document)
		
		befExp := docT.ExpiresAt
		docT.Update(value, expAfterSec)
		
		c.OpLock.Lock()
		c.putOrUpdateTtlIndex(docT, befExp)
		defer c.OpLock.Unlock()
		return docT
		
	} else {
		return c.Put(key, value, expAfterSec)
	}
}


func (c *Collection) putOrUpdateTtlIndex(
		doc *document.Document,
		beforeExpires interface{}) {
	c.checkOpened()
	
	// b(o) -> a(o) = update index
	if beforeExpires != nil && doc.ExpiresAt != nil {
		beat := beforeExpires.(int64)
		deat := doc.ExpiresAt.(int64)
		
		if b, exists := c.TtlIndex.Get(beat); exists {
			b.(*Bucket).Remove(doc.Key)
		}
		
		if b, exists := c.TtlIndex.Get(deat); exists {
			b.(*Bucket).Add(doc.Key)
		} else {
			bucket := NewBucket(deat)
			bucket.Add(doc.Key)
			c.TtlIndex.Put(deat, bucket)
		}
		
	// b(o) -> a(x) = remove index
	} else if beforeExpires != nil && doc.ExpiresAt == nil {
		eat := beforeExpires.(int64)
		if b, exists := c.TtlIndex.Get(eat); exists {
			b.(*Bucket).Remove(doc.Key)
		}
		
	// b(x) -> a(o) = add index
	} else if beforeExpires == nil && doc.ExpiresAt != nil {
		eat := doc.ExpiresAt.(int64)
		if b, exists := c.TtlIndex.Get(eat); exists {
			b.(*Bucket).Add(doc.Key)
		} else {
			bucket := NewBucket(eat)
			bucket.Add(doc.Key)
			c.TtlIndex.Put(eat, bucket)
		}
	}
}


func (c *Collection) Get(key string) (interface{}, bool) {
	c.checkOpened()
	return c.Documents.Get(key)
}


func (c *Collection) Remove(key string) (bool) {
	c.checkOpened()
	
	if doc, exists := c.Documents.Get(key); exists {
		// Remove index
		c.OpLock.Lock()
		docT := doc.(*document.Document);
		if docT.ExpiresAt != nil {
			if bucket, exists := c.TtlIndex.Get(docT.ExpiresAt.(int64)); exists {
				bucket.(*Bucket).Remove(key)
			}
		}
		
		// Remove document
		c.Documents.Remove(key)
		defer c.OpLock.Unlock()
		return true
	}
	
	return false
}


func (c *Collection) Size() (int) {
	return c.Documents.Size()
}


func (c *Collection) Clear() {
	c.checkOpened()
	c.OpLock.Lock()
	
	// Clear TTL indexes
	c.TtlIndex.Clear()
	// Clear Documents
	c.Documents.Clear()
	
	defer c.OpLock.Unlock()
}


func (c *Collection) ContainsKey(key string) (contains bool) {
	c.checkOpened()
	c.OpLock.Lock()
	_, contains = c.Documents.Get(key)
	defer c.OpLock.Unlock()
	
	return
}


func (c *Collection) Close() (*Collection) {
	c.checkOpened()
	c.OpLock.Lock()
	
	// Update status
	c.Opened = false
	
	// Close all expiration handlers
	for _, handler := range c.expirationHandler {
		handler.Close()
	}
	
	// Stop all threads
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
	}()
	wg.Wait()
	//
	
	defer c.OpLock.Unlock()
	
	log.Printf("Collection [ %s ] closed", c.Name)
	return c
}
