package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router.Use(middleware.Logger)
	buildRoutes(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Config.FileServer.Port), router))
}

// buildRoots sets up all routes from the configuration file
func buildRoutes(router chi.Router) {
	for _, file := range Config.Files {
		filename := filepath.Base(file.Endpoint)
		router.Get(file.Endpoint, func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
			http.ServeFile(rw, r, file.FilePath)
		})
	}
}
