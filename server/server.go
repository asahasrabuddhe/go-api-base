package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	db "gitlab.com/asahasrabuddhe/go-api-base/database"
	"gitlab.com/asahasrabuddhe/go-api-base/router"
	"gitlab.com/asahasrabuddhe/go-api-base/server/middleware"
	"log"
	"net/http"
)

func Init(path, filename string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("json")
	//viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to read config", err)
	}

	router.Router = mux.NewRouter()
	router.ApiRouter = router.Router.PathPrefix("/" + viper.GetString("app.api_prefix")).Subrouter()

	router.Router.Use(middleware.FilterIncomingRequestsMiddleware)

	credentials := viper.GetStringMapString("database.default")
	db.Open(credentials["username"], credentials["password"], credentials["name"], credentials["host"], credentials["port"])
}

func Start() error {
	allowedHeaders := handlers.AllowedHeaders(viper.GetStringSlice("app.cors.allowed_headers"))
	allowedOrigins := handlers.AllowedOrigins(viper.GetStringSlice("app.cors.allowed_origins"))
	allowedMethods := handlers.AllowedMethods(viper.GetStringSlice("app.cors.allowed_methods"))
	// start server listen
	router.Router.PathPrefix("/images/").
		Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(viper.GetString("public_path")))))
	// with error handling
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", viper.GetString("app.address"), viper.GetString("app.port")), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router.Router))
	return err
}
