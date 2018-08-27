package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	db "gitlab.com/asahasrabuddhe/go-api-base/database"
	"gitlab.com/asahasrabuddhe/go-api-base/router"
	"gitlab.com/asahasrabuddhe/go-api-base/server/middleware"
	"net/http"
)

func Init(username, password, database, host, port string) {
	router.Router = mux.NewRouter()

	router.Router.Use(middleware.FilterIncomingRequestsMiddleware)

	db.Open(username, password, database, host, port)
}

func Start(address string, port string) error {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	// start server listen
	router.Router.PathPrefix("/images/").
		Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("/home/ajitem/Projects/safebaby/bin/public/images"))))
	// with error handling
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router.Router))
	return err
}
