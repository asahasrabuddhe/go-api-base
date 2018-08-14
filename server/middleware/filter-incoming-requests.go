package middleware

import (
	"encoding/json"
	"gitlab.com/asahasrabuddhe/go-api-base/response"
	"net/http"
)

func FilterIncomingRequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.Header.Get("Content-Type") != "application/json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := response.Response{
				Success: false,
				Message: "API only accepts JSON input. Please format the input properly or ensure that correct headers are set in the request header.",
			}
			encoder := json.NewEncoder(w)
			encoder.Encode(&res)
		}
		next.ServeHTTP(w, r)
	})
}
