package middleware

import (
	"net/http"
)

// Limit simultaneous requests to handler
func LimitRequests(next http.HandlerFunc, limit int32) http.HandlerFunc {
	requestsSemaphore := make(chan struct{}, limit)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// When request handling ended free semaphore
		defer func() {
			<-requestsSemaphore
		}()
		requestsSemaphore <- struct{}{}
		next(w, r)
	})
}
