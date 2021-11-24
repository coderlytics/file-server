package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

// Start the server with the given configuration
func Start(configPath string) {
	if err := initConfig(configPath); err != nil {
		log.Fatal(err)
	}
	initLogging()
	run()
}

// initLogging initializes the logging subsystem with the configured log level
func initLogging() {
	lvl, err := log.ParseLevel(Config.Logging.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(lvl)

	if log.GetLevel() == log.TraceLevel {
		log.SetReportCaller(true)
	}
}

// run starts the actual server
func run() {
	router := chi.NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Config.FileServer.Port), router))
}
