package middlewares

import (
	"log"
	"net/http"
	"time"
)

type HandlerSign func(w http.ResponseWriter, r *http.Request)

//Log ..
func Log(f HandlerSign) HandlerSign {
	return func(w http.ResponseWriter, r *http.Request) {
		defer TimeTrack(time.Now(), r.URL.Path)
		log.Printf("request %q, method %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

// Auth validation
func Auth(f HandlerSign) HandlerSign {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "token-is-secure" {
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No Authorization"))
}

// TimeTrack verify the time
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%q execution time: %s", name, elapsed)
}