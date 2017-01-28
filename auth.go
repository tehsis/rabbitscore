package main

// // import (
// // 	"log"
// // 	"net/http"
// // 	"time"
// // )

// // // Logger is a decorator function to log requests
// // func Logger(inner http.Handler, name string) http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 		start := time.Now()

// // 		inner.ServeHTTP(w, r)

// // 		log.Printf(
// // 			"%s\t%s\t%s\t%s",
// // 			r.Method,
// // 			r.RequestURI,
// // 			name,
// // 			time.Since(start),
// // 		)
// // 	})
// // }

// package main

// import "net/http"
// import JWS "github.com/SermoDigital/jose/JWS"

// // Auth is a decorator function to parse auth headers
// func Auth(inner http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		inner.ServeHTTP(w, r)
// 		jwt, err := JWS.ParseJWTFromRequest(r)
// 	})
// }
