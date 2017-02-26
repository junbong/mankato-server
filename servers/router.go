package server

import (
	"github.com/Junbong/mankato-server/configs"
	"github.com/Junbong/mankato-server/db/collections"
	"github.com/Junbong/mankato-server/db/documents"
	"github.com/Junbong/mankato-server/utils"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
)

type ServerRouter struct {
	Configuration *configs.Config
	Mux           *mux.Router
	Database      *collection.Collection
}

type ServerInfo struct {
	Status     string `json:"status"`
	ServerAddr string `json:"server_addr"`
	Keys       int `json:"keys"`
}

var serverinf *ServerInfo


func New(config *configs.Config, database *collection.Collection) *ServerRouter {
	return &ServerRouter{Configuration:config, Mux:mux.NewRouter(), Database:database}
}


func (r *ServerRouter) BeginRoutes() {
	// Basic index
	r.Mux.HandleFunc("/", r.Index).Methods("GET")
	
	// Server
	r.Mux.HandleFunc("/_info", r.GetServer).Methods("GET")
	
	// Document
	r.Mux.HandleFunc("/{key}", r.GetDocument).Methods("GET")                // Get document with specified key
	r.Mux.HandleFunc("/{key}", r.CreateOrUpdateDocument).Methods("POST")    // Create new document with specified key
	r.Mux.HandleFunc("/{key}", r.DeleteDocument).Methods("DELETE")          // Remove document with specified key
	
	// Start
	serverinf = &ServerInfo{
		Status: "ok",
		ServerAddr: fmt.Sprintf("%s:%d", r.Configuration.Server.Host, r. Configuration.Server.Port),
	}
	log.Printf("Running with... %s\n", serverinf.ServerAddr)
	log.Fatal(http.ListenAndServe(serverinf.ServerAddr, r.Mux))
}


func onPreRequest(r *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
	}
}


func (sr *ServerRouter) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}


func (sr *ServerRouter) GetServer(
		w http.ResponseWriter,
		r *http.Request) {
	onPreRequest(r)
	
	log.Println(sr.Database.String())
	
	serverinf.Keys = sr.Database.Size()
	onResultJson(w, serverinf)
}


func (sr *ServerRouter) GetDocument(
		w http.ResponseWriter,
		r *http.Request) {
	onPreRequest(r)
	vars := mux.Vars(r)
	keyOfData := vars["key"]
	
	doc, exists := sr.Database.Get(keyOfData)
	
	// TODO: make result to JSON with lib
	if exists {
		docT := doc.(*document.Document)
		onResultJson(w, docT)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}


func (sr *ServerRouter) CreateOrUpdateDocument(
		w http.ResponseWriter,
		r *http.Request) {
	onPreRequest(r)
	vars := mux.Vars(r)
	keyOfData := vars["key"]
	
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		var valueOfData string
		
		if len(b) > 0 {
			valueOfData = string(b)
		} else {
			valueOfData = ""
		}
		
		// TTL option
		ttl := r.FormValue("ttl")
		var expAfterSec int = -1
		
		if utils.IsNotNilOrEmpty(ttl) {
			if i, e := strconv.Atoi(ttl); e == nil {
				expAfterSec = i
			}
		}
		
		doc := sr.Database.PutOrUpdate(keyOfData, valueOfData, expAfterSec)
		onResultJson(w, doc)
		
	} else {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}


func (sr *ServerRouter) DeleteDocument(
		w http.ResponseWriter,
		r *http.Request) {
	onPreRequest(r)
	vars := mux.Vars(r)
	keyOfData := vars["key"]
	
	if sr.Database.Remove(keyOfData) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}


func (sr *ServerRouter) NotSupported(
		w http.ResponseWriter,
		r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "This operation does not supported")
}
