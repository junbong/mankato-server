package db

import "fmt"

type Database struct {
	buckets map[string]Bucket
}

type Bucket struct {
	name    string
	data    map[string]Data
	expire  uint16
}

type Data struct {
	key     string
	value   string
	expire  uint16
}


var database Database


func New() {
	database := new(Database)
	fmt.Println("database", database)
	fmt.Println("&database", &database)
	fmt.Println("*database", *database)
}


func Put(bucket, key, value string) {
	Bucket, exists := database.buckets[bucket]
	
	// Create new bucket when not exists
	if !exists {
		Bucket = createBucket(bucket, 0)
	}
	
	Bucket.data[key] = Data{key:key, value:value}
	
	fmt.Println("Database", database)
}


func createBucket(name string, expire uint16) Bucket {
	bucket := Bucket{name:name, expire:expire}
	//bucket.data["nono"] = Data{}
	fmt.Println("New Bucket", bucket)
	
	database.buckets[name] = bucket
	return bucket
}
