package main

import (
	"github.com/Junbong/mankato-server/db/collections"
	"github.com/Junbong/mankato-server/db/documents"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)


var Collection *collection.Collection

func BeginRoutes(col *collection.Collection, host string, port int) {
	// TODO: naive reference
	Collection = col
	
	// Primary routing points started from here
	router := httprouter.New()
	
	// Basic index
	router.GET("/", Index)
	
	// Document
	router.GET("/:key", GetDocument)                // Get document with specified key
	router.POST("/:key", CreateDocument)            // Create new document with specified key
	router.DELETE("/:key", DeleteDocument)          // Remove document with specified key
	
	// Meta Information
	//router.GET("/_info", GetServer)      // Get overall server information
	
	// Listen requests
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router))
}


func onPreRequest(
		r *http.Request,
		params httprouter.Params) {
	log.Printf("[%s] %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
}


func onResultJson(
		w http.ResponseWriter,
		obj interface{}) {
	// TODO: memoize result of marshaling
	if b, err := json.Marshal(obj); err == nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, string(b[:]))
	} else {
		log.Fatal(err)
		w.WriteHeader(500)
	}
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


func GetDocument(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	keyOfData := params.ByName("key")
	
	doc, exists := Collection.Get(keyOfData)
	
	// TODO: make result to JSON with lib
	if exists {
		docT := doc.(*document.Document)
		onResultJson(w, docT)
	} else {
		w.WriteHeader(404)
	}
}


func CreateDocument(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	keyOfData := params.ByName("key")
	
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		var valueOfData string
		
		if len(b) > 0 {
			valueOfData = string(b)
		} else {
			valueOfData = ""
		}
		
		// TODO: retrieve query params
		var expireOfData uint = 0
		
		Collection.Put(keyOfData, valueOfData, expireOfData)
		w.WriteHeader(200)
	} else {
		log.Fatal(err)
		w.WriteHeader(500)
	}
}


func DeleteDocument(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	onPreRequest(r, params)
	keyOfData := params.ByName("key")
	
	if Collection.Remove(keyOfData) {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
}


func NotSupported(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(404)
	fmt.Fprint(w, "This operation does not supported!\n")
}
