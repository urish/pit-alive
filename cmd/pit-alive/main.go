package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/host"
)

const port = 8080

func main() {
	fmt.Println("Starting HTTP server on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		remote := r.RemoteAddr
		forwardedFor := r.Header.Get("X-Forwarded-For")
		if forwardedFor != "" {
			remote = forwardedFor
		}

		fmt.Printf("[%s] %s %s %s\n", time.Now().Format(time.RFC3339), remote, r.Method, r.URL)
		uptime, _ := host.Uptime()
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(
			"{\"alive\": true, \"uptime\": %d}", uptime,
		)))
	})

	/* Start the HTTP server:*/
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
