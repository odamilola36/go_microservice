package middlewares

import (
	"log"
	"microservices/security"
	"net/http"
	"strings"
	"time"
)

func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		log.Printf(
			`{"proto": "%s", "route" :"%s%s", "method": "%s", "time": "%s"}`,
			r.Proto, r.Host, r.RequestURI, r.Method, time.Since(t),
		)
	}
}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractToken, err2 := security.ExtractToken(r)
		if err2 != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimSpace(extractToken)
		toke, err := security.ParseToke(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !toke.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
