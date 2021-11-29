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
	log.Info(fmt.Sprintf("Starting server on port %s", Config.FileServer.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Config.FileServer.Port), router))
}

// buildRoots sets up all routes from the configuration file
func buildRoutes(router chi.Router) {
	for _, file := range Config.Files {
		log.Debugf("Adding endpoint %s for file %s", file.Endpoint, file.FilePath)

		router.Get(file.Endpoint, func(file File) func(rw http.ResponseWriter, r *http.Request) {
			return func(rw http.ResponseWriter, r *http.Request) {
				// if token is set, authorization check will be performed
				if len(file.Token) > 0 {
					log.Debugf("Adding authentication to endpoint %s", file.Endpoint)
					auth := r.Header.Get("Authorization")

					if file.Token != auth {
						http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
						return
					}
				}

				rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(file.Endpoint)))
				http.ServeFile(rw, r, file.FilePath)
			}
		}(file))
	}
}
