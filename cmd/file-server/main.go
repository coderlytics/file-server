package fileserver

import (
	"flag"

	"coderlytics.io/file-server/internal/server"
)

func main() {
	var (
		configPath = flag.String("config", "config.yml", "Path to the config file")
	)

	flag.Parse()
	server.Start(*configPath)
}
