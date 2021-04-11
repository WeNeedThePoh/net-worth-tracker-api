package http

import (
	"encoding/json"
	"net/http"
)

func handleCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Access-Control-Request-Method, X-Requested-With, X-Authorization, X-Locale")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func handleRouteNotFound(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Route requested not found",
			})
			return
		}
	})
}
