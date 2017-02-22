package main

import (
	"log"
	"net/http"
	"fmt"
	"reflect"
	"github.com/Junbong/mankato-server/db"
	"github.com/julienschmidt/httprouter"
)


var database db.Database

func BeginRoutes(lDatabase *db.Database, host string, port int) {
	// TODO: naive reference
	database = *lDatabase
	
	// Primary routing points started from here
	router := httprouter.New()
	
	// Basic index
	router.GET("/", Index)
	
	// Bucket
	//router.GET("/:collection/_INFO", NotSupported)    // Get information of specified collection
	router.GET("/:collection", GetCollection)           // Get all key & values in specified collection
	router.POST("/:collection", CreateCollection)       // Create new collection with specified name
	router.DELETE("/:collection", NotSupported)         // Delete collection, all keys & values in specified collection
	
	// Key
	router.GET("/:collection/:key", NotSupported)
	router.POST("/:collection/:key", NotSupported)
	router.POST("/:collection/:key/:value", NotSupported)
	router.DELETE("/:collection/:key", NotSupported)
	
	// Listen requests
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router))
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


func GetCollection(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", params.ByName("collection"))
}


func CreateCollection(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	nameOfCollection := params.ByName("collection")
	log.Printf("[%s] %s", r.Method, r.URL)
	
	var collection *db.Collection
	
	collection, _ = database.GetCollection(nameOfCollection, true)
	log.Println("TypeOf", reflect.TypeOf(collection))
	
	// TODO: make collection interface
	// TODO: make result to JSON with lib
	fmt.Fprint(w, fmt.Sprintf("{ \"collection\": { \"name\": \"%s\", \"size\": %d } }", collection, 1))
}


func NotSupported(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(404)
	fmt.Fprint(w, "This operation does not supported!\n")
}
