package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/Junbong/mankato-server/db"
)

const (
	version = "0.0.1"
)


var host = flag.String("h", "localhost", "Set host address of server")
var port = flag.Int("p", 7120, "Set port number of server")


func main() {
	// Parse program flags and read configurations
	flag.Parse()
	log.Println("Project Mankato", version)
	log.Printf("Running with... %s:%d\n", *host, *port)
	
	// Initialize database
	database := db.New()
	log.Println("New database", database)
	log.Println("New database", *database)
	
	// TODO: Tests
	//database.Put("test", "foo", "bar")
	//database.Put("test", "foo1", "bar1")
	//database.Put("newday", "has", "come")
	//collection, _ := database.GetCollection("test", false)
	//nilCollection, _ := database.GetCollection("not_exists", false)
	//log.Println("GetCollection", "test", collection)
	//log.Println("GetCollection", "not_exists", nilCollection)
	
	// Start router & server
	//watchSysSigs(shutdownGraceful)
	
	// Watch system signal
	BeginRoutes(database, *host, *port)
}


func watchSysSigs(termination func()) {
	sigs := make(chan os.Signal, 1)
	term := make(chan bool, 1)
	
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		sig := <-sigs
		fmt.Println()
		log.Println(sig)
		term <-true
	}()
	
	<-term
	termination()
}


func shutdownGraceful() {
	log.Println("Shutdown graceful...")
	
	// Shutdown
	
	log.Println("Bye :]")
}
