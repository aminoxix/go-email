package configs

import (
	"net/http"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow requests from the detected host
		// get referer and ensure it's not empty
		referer := r.Referer()
		if len(referer) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", referer[:len(referer)-1])
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*") // fallback to wildcard or a fixed domain
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
