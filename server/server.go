package server

import (
	"fmt"
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
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), router.Router)
	return err
}
