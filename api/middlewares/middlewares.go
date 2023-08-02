package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		log.Printf(`method: "%s", route: "%s%s", request_time: "%v"`, r.Method, r.Host, r.URL.Path, time.Since(startTime))
	}
}
