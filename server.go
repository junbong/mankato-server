package main

import (
	"github.com/Junbong/mankato-server/servers"
	"github.com/Junbong/mankato-server/db/collections"
	"flag"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	
	// Initialize configs
	con := &server.Config{Host:*host, Port:*port}
	
	// Initialize database
	col := collection.New("default")
	col.Open()
	
	// Watch system signal
	//watchSysSigs(shutdownGraceful)
	
	// Start router & server
	svr := server.New(con, col)
	svr.BeginRoutes()
	
	defer col.Close()
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


func shutdownGraceful(col *collection.Collection) {
	log.Println("Shutdown graceful...")
	
	// Shutdown
	col.Close()
	
	log.Println("Bye :]")
}
