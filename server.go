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
	"sync"
)

const (
	version = "0.0.1"
)


var (
	ophost = flag.String("h", "localhost", "Set host address of server")
	opport = flag.Int("p", 7120, "Set port number of server")
	opconf = flag.String("conf", "./etc/configuration.yml", "Configuration path")
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
	col := collection.New(config.Collection.DefaultName, config)
	col.Open()
	
	// Start router & server
	svr := server.NewRouter(config, col)
	svr.SetupRoutes()
	
	// Run server and watch system signal
	watchSysSigs(svr.StartServe, svr.StopServe)
	
	// Called when shutdown
	defer shutdownGraceful(col)
}


func watchSysSigs(fn1, fn2 func()) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	
	var wg sync.WaitGroup
	wg.Add(1)
	
	go func() {
		defer wg.Done()
		fn1()
	}()
	
	select {
	case sig := <-sigs:
		fmt.Println()
		log.Println(sig, "signal")
	}
	
	fn2()
	wg.Wait()
}


func shutdownGraceful(col *collection.Collection) {
	log.Println("Shutdown graceful...")
	
	// Shutdown
	col.Close()
	
	log.Println("Bye :]")
}
