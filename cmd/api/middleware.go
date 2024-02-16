package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "CloseS")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func (app *application) rateLimit(next http.Handler) http.Handler {
	lateLimiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !lateLimiter.Allow() {
			app.rateLimitExceedResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
