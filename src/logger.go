package main

import (
    "log"
    "net/http"
    "time"
)

// log the request method, the request uniform resource identifier, 
// the name..., and the time since start...
func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
    )

    })
}
