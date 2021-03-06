package server

import (
	"fmt"
	db "github.com/asahasrabuddhe/go-api-base/database"
	"github.com/asahasrabuddhe/go-api-base/mail"
	"github.com/asahasrabuddhe/go-api-base/middleware"
	"github.com/asahasrabuddhe/go-api-base/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

	mail.BootstrapMail()

	credentials := viper.GetStringMapString("database.default")
	db.Open(credentials["username"], credentials["password"], credentials["name"], credentials["host"], credentials["port"])
}

func Start() (err error) {
	allowedHeaders := handlers.AllowedHeaders(viper.GetStringSlice("app.cors.allowed_headers"))
	allowedOrigins := handlers.AllowedOrigins(viper.GetStringSlice("app.cors.allowed_origins"))
	allowedMethods := handlers.AllowedMethods(viper.GetStringSlice("app.cors.allowed_methods"))
	// start server listen
	router.Router.PathPrefix("/public/").
		Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(viper.GetString("public_path")))))
	// with error handling
	log.Printf("Server Started On IP - %v PORT - %v", viper.GetString("app.address"), viper.GetString("app.port"))
	if viper.GetBool("app.tls") {
		log.Println(http.ListenAndServeTLS(fmt.Sprintf("%v:%v", viper.GetString("app.address"), viper.GetString("app.port")), viper.GetString("app.cert_path"), viper.GetString("app.key_path"), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router.Router)))
	} else {
		log.Println(http.ListenAndServe(fmt.Sprintf("%v:%v", viper.GetString("app.address"), viper.GetString("app.port")), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router.Router)))
	}
	return
}
