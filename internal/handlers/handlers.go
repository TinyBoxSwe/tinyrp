package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	config "tinyrp/internal/configs"
	logger "tinyrp/internal/loggers"
)

func ProxyRequestHandler(routes map[string]config.Resource) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.LogIncomingRequest(request)

		targetURL := request.URL

		parts := strings.Split(targetURL.Path, "/")
		if len(parts) < 2 {
			http.Error(writer, "Invalid request", http.StatusBadRequest)
			logger.LogError(fmt.Errorf("invalid request: URL path too short: %s", targetURL.Path))
			return
		}

		target_server := parts[1]

		route, ok := routes[target_server]
		if !ok {
			http.Error(writer, "Unknown route", http.StatusNotFound)
			logger.LogError(fmt.Errorf("unknown route: %s", target_server))
			return
		}

		destinationURL := route.Destination_url

		forwardRequest, err := http.NewRequest(request.Method, destinationURL, request.Body)
		if err != nil {
			http.Error(writer, "Could not create request", http.StatusInternalServerError)
			logger.LogError(fmt.Errorf("could not create forward request to %s: %v", destinationURL, err))
			return
		}
		defer request.Body.Close()

		for key, values := range request.Header {
			for _, value := range values {
				forwardRequest.Header.Add(key, value)
			}
		}

		client := &http.Client{}
		resp, err := client.Do(forwardRequest)
		if err != nil {
			http.Error(writer, "Error forwarding request", http.StatusBadGateway)
			logger.LogError(fmt.Errorf("error forwarding request to %s: %v", destinationURL, err))
			return
		}
		defer resp.Body.Close()

		writer.WriteHeader(resp.StatusCode)
		for key, values := range resp.Header {
			for _, value := range values {
				writer.Header().Add(key, value)
			}
		}

		if _, err = io.Copy(writer, resp.Body); err != nil {
			http.Error(writer, "Error writing response body", http.StatusInternalServerError)
			logger.LogError(fmt.Errorf("error writing response body: %v", err))
		}
	}
}
