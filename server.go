package main

import (
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
	log.Printf("Running with... %s:%d\n", *host, *port)
	
	// Initialize cache
	
	// Start router & server
	
	// Watch system signal
	watchSysSigs(shutdownGraceful)
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
