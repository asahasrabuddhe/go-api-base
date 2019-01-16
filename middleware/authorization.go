package middleware

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/asahasrabuddhe/go-api-base/database"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"log"
	"net/http"
	"strings"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "{'status': false, 'message': 'Please send auth token with the request'}")
			return
		} else {
			auth := r.Header.Get("Authorization")
			token := strings.Split(auth, " ")
			if len(token) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "{'status': false, 'message': 'Please send auth token with the request'}")
			} else if len(token) == 2 && token[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "{'status': false, 'message': 'Please send auth token with the request'}")
			} else {
				jwtToken, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
						if token.Method.Alg() != "HS256" {
							return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
						}
					}
					return []byte("safebaby"), nil
				})

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					log.Println(err)
					fmt.Fprintln(w, "{'status': false, 'message': 'Please send auth token with the request'}")
					return
				}

				if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
					var user_id int

					row := database.DB.QueryRow("SELECT user_id FROM user_auth_tokens WHERE token = ? AND deleted_on IS NULL", claims["jti"])

					if err := row.Scan(&user_id); err != nil {
						fmt.Fprintln(w, "{'status': false, 'message': 'Token Expired'}")
						return
					} else if float64(user_id) != claims["aud"].(float64) {
						fmt.Fprintln(w, "{'status': false, 'message': 'Invalid Token'}")
						return
					} else {
						context.Set(r, "id", claims["aud"])
						context.Set(r, "role", claims["rle"])
					}
				} else {
					fmt.Println(err)
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
