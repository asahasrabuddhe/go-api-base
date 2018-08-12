package middleware

import (
	"fmt"
	"net/http"
)

func FilterIncomingRequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "{'status': 'failure', "+
				"'error': 'API only accepts JSON input. Please format the input properly or ensure that correct headers are set in the request header.'}")
			return
		}
		next.ServeHTTP(w, r)
	})
}
