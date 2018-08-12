package middleware

import (
	"fmt"
	"github.com/gorilla/context"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"log"
	"net/http"
	"strings"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "{'status': 'failure', 'error': 'Please send auth token with the request'}")
			return
		} else {
			auth := r.Header.Get("Authorization")
			token := strings.Split(auth, " ")
			if len(token) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "{'status': 'failure', 'error': 'Please send auth token with the request'}")
			} else if len(token) == 2 && token[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "{'status': 'failure', 'error': 'Please send auth token with the request'}")
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
					fmt.Fprintln(w, "{'status': 'failure', 'error': 'Please send auth token with the request'}")
					return
				}

				if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
					context.Set(r, "id", claims["aud"])
					context.Set(r, "role", claims["role"])
				} else {
					fmt.Println(err)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
