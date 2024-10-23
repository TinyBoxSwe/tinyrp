package main

import (
	"fmt"
	"net/http"
	"os"
	config "tinyrp/internal/configs"
	handler "tinyrp/internal/handlers"
	logger "tinyrp/internal/loggers"
)

func RunServer(cfg config.Configuration) {
	address := cfg.Server.Host + ":" + cfg.Server.Listen_port
	fmt.Fprintf(os.Stdout, "Server listening on %s\n", address)

	endpointMap := make(map[string]config.Resource)
	for _, resource := range cfg.Resources {
		endpointMap[resource.Name] = resource
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.ProxyRequestHandler(endpointMap))

	if err := http.ListenAndServe(address, mux); err != nil {
		logger.LogError(fmt.Errorf("server error: %v", err))
	}
}

func main() {
	config, err := config.Load()
	if err != nil {
		logger.LogError(fmt.Errorf("error loading config: %v", err))
		os.Exit(1)
	}

	RunServer(config)
}
