package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"fmt"
)


func BeginRoute(host string, port int) {
	// Primary routing points started from here
	router := httprouter.New()
	
	// Basic index
	router.GET("/", Index)
	
	// Bucket
	//router.GET("/:bucket/_INFO", NotSupported)  // Get information of specified bucket
	router.GET("/:bucket", GetBucket)           // Get all key & values in specified bucket
	router.POST("/:bucket", CreateBucket)       // Create new bucket with specified name
	router.DELETE("/:bucket", NotSupported)     // Delete bucket, all keys & values in specified bucket
	
	// Key
	router.GET("/:bucket/:key", NotSupported)
	router.POST("/:bucket/:key", NotSupported)
	router.POST("/:bucket/:key/:value", NotSupported)
	router.DELETE("/:bucket/:key", NotSupported)
	
	// Listen requests
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router))
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


func GetBucket(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", params.ByName("bucket"))
}


func CreateBucket(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "New bucket created: %s\n", params.ByName("bucket"))
}


func NotSupported(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(404)
	fmt.Fprint(w, "This operation does not supported!\n")
}
