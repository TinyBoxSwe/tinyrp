package loggers

import (
	"fmt"
	"net/http"
	"time"
)

func LogIncomingRequest(request *http.Request) {
	time := time.Now().UTC().Format("2006-01-02 15:04:05.000")
	path := request.URL.Path
	method := request.Method

	fmt.Printf("[%s] %s %s\n", time, method, path)
}

func LogError(err error) {
	time := time.Now().UTC().Format("2006-01-02 15:04:05.000")
	fmt.Printf("[%s] ERROR: %v\n", time, err)
}
