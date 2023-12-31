package middleware

import (
	"log"
	"net/http"

	"github.com/rcgc/person-api-http-net/authorization"
)

// Log .
func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("petición %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

// Authentication .
func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			// responder "prohibido"
			forbidden(w, r)
			return
		}

		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}