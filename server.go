package main

import (
	"github.com/Junbong/mankato-server/servers"
	"github.com/Junbong/mankato-server/configs"
	"github.com/Junbong/mankato-server/db/collections"
	"gopkg.in/yaml.v2"
	"flag"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
)

const (
	version = "0.0.1"
)


var (
	ophost = flag.String("h", "localhost", "Set host address of server")
	opport = flag.Int("p", 7120, "Set port number of server")
	opconf = flag.String("conf", "./conf/default.yml", "Configuration path")
	config  *configs.Config
)


func init() {
	log.Println("Project Mankato", version)
	
	// Parse program flags and read configurations
	flag.Parse()
	
	// Load configuration
	if dat, err := ioutil.ReadFile(*opconf); err == nil {
		config = &configs.Config{}
		if err := yaml.Unmarshal([]byte(dat), config); err == nil {
			// Use host command option first
			if config.Server.Host != *ophost {
				config.Server.Host = *ophost
			}
			// Use port command option first
			if config.Server.Port != *opport {
				config.Server.Port = *opport
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
	log.Printf("Use configuration %s", *opconf)
}


func main() {
	// Initialize database
	col := collection.New("default", config)
	col.Open()
	
	// Watch system signal
	//watchSysSigs(shutdownGraceful)
	
	// Start router & server
	svr := server.New(config, col)
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
