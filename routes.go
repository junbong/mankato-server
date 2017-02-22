package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/Junbong/mankato-server/db/database"
	"github.com/julienschmidt/httprouter"
	"github.com/Junbong/mankato-server/db/collection"
	"github.com/Junbong/mankato-server/db/data"
)


var db *database.Database

func BeginRoutes(lDb *database.Database, host string, port int) {
	// TODO: naive reference
	db = lDb
	
	// Primary routing points started from here
	router := httprouter.New()
	
	// Basic index
	router.GET("/", Index)
	
	// Server
	router.GET("/info", GetServer)      // Get overall server information
	
	// Collections
	router.GET("/collections/:collection", GetCollection)           // Get collection with specified name
	router.POST("/collections/:collection", CreateCollection)       // Create new collection with specified name
	router.DELETE("/collections/:collection", DeleteCollection)     // Delete collection, all keys & values in specified collection
	
	// Data
	router.GET("/collections/:collection/:key", GetData)                // Get key-value data with specified key
	router.POST("/collections/:collection/:key", CreateData)            // Create new key-value data
	router.POST("/collections/:collection/:key/:value", NotSupported)
	router.DELETE("/collections/:collection/:key", NotSupported)
	
	// Listen requests
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router))
}


func onPreRequest(
		r *http.Request,
		params httprouter.Params) {
	log.Printf("[%s] %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


func GetServer(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	
	// TODO: server references
	fmt.Fprint(w, fmt.Sprintf("{ \"server\": { \"host\": \"%s\", \"port\": %d, \"collections\": [] } }", "HOST_ADDR", 7120))
}


func GetCollection(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	nameOfCollection := params.ByName("collection")
	
	col, exists := db.Get(nameOfCollection)
	
	// TODO: make result to JSON with lib
	if exists {
		colT := col.(*collection.Collection)
		fmt.Fprint(w, fmt.Sprintf("{ \"collection\": { \"name\": \"%s\", \"size\": %d } }", colT.Name(), colT.Size()))
	} else {
		w.WriteHeader(404)
	}
}


func CreateCollection(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	nameOfCollection := params.ByName("collection")
	
	// TODO: retrieve query params

	col := db.GetOrCreateCollection(nameOfCollection, true).(*collection.Collection)

	// TODO: make result to JSON with lib
	fmt.Fprint(w, fmt.Sprintf("{ \"collection\": { \"name\": \"%s\", \"size\": %d } }", col.Name(), col.Size()))
}


func DeleteCollection(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	nameOfCollection := params.ByName("collection")
	
	_, exists := db.Get(nameOfCollection)
	
	if exists {
		db.Remove(nameOfCollection)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
}


func GetData(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	nameOfCollection := params.ByName("collection")
	keyOfData := params.ByName("key")
	
	col, exists := db.Get(nameOfCollection)
	
	// TODO: make result to JSON with lib
	if exists {
		colT := col.(*collection.Collection)
		
		d, dexists := colT.Get(keyOfData)
		
		if dexists {
			dT := d.(*data.Data)
			fmt.Fprint(w, fmt.Sprintf("{ \"data\": { \"key\": \"%s\", \"value\": \"%s\", \"expire\": %d } }", dT.Key(), dT.Value(), dT.Expire()))
		} else {
			w.WriteHeader(404)
		}
	} else {
		w.WriteHeader(404)
	}
}


func CreateData(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	nameOfCollection := params.ByName("collection")
	keyOfData := params.ByName("key")
	valueOfData := "parse_value_at_here"
	expireOfData := 0
	
	// TODO: retrieve query params
	
	col := db.GetOrCreateCollection(nameOfCollection, true).(*collection.Collection)
	col.Put(keyOfData, valueOfData, expireOfData)
	fmt.Fprint(w, fmt.Sprintf("{ \"collection\": { \"name\": \"%s\", \"size\": %d } }", col.Name(), col.Size()))
}


func NotSupported(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(404)
	fmt.Fprint(w, "This operation does not supported!\n")
}
