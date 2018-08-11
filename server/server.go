package server

import (
	"net/http"
	"fmt"
	"gitlab.com/asahasrabuddhe/go-api-base/router"
	"github.com/gorilla/mux"
	db "gitlab.com/asahasrabuddhe/go-api-base/database"
)

func Init(username, password, database, host, port string) {
	router.Router = mux.NewRouter()
	db.Open(username, password, database, host, port)
}

func Start(address string, port string) error {
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), router.Router)
	return err
}
