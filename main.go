package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dronemill/eventsocket"
)

var stopped chan bool

func main() {
	// parse cli flags
	flag.Parse()

	// if we are just supposed to print the version, then do so
	if printVersion {
		fmt.Printf("maestro %s\n", Version)
		os.Exit(0)
	}

	// load configuration
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Maestro is starting...")

	log.WithField("port", config.Eventsocket.Port).WithField("maxMessage", 5242880).Info("Creating new server")
	es, err := eventsocket.NewServer(fmt.Sprintf(":%d", config.Eventsocket.Port))
	if err != nil {
		log.Fatal(err.Error())
	}

	es.SetDefaultMaxMessageSize(5242880) // 5MB

	log.WithField("port", config.Eventsocket.Port).Info("Starting server")
	go es.Start()

	stopped = make(chan bool)
	<-stopped

	log.Info("Maestro is shutting down...")
}
