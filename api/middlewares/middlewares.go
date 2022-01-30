package middlewares

import (
	"errors"
	"flag"
	"net/http"

	"MNZ/api/auth"
	"MNZ/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	insecure := flag.Bool("insecure", false, "Set insecure mode to bypass JWT Auth")
	flag.Parse()

	if *insecure != true {
		return func(w http.ResponseWriter, r *http.Request) {
			err := auth.TokenValid(r)
			if err != nil {
				responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
				return
			}
			next(w, r)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
