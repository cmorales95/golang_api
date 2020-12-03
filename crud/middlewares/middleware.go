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

//TimeTrack verify the time
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%q execution time: %s", name, elapsed)
}