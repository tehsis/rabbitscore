package main

import "net/http"

// SetContentType is a decorator function to set Content-type
func SetContentType(inner http.Handler, contentType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		inner.ServeHTTP(w, r)
	})
}

func SetAccessControl(inner http.Handler, ac string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", ac)
		inner.ServeHTTP(w, r)
	})
}
