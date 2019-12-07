package bootstrap

import (
	"github.com/Vespen/steam-openid-redirect/server"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
)

// Creates server.
func NewServer() server.Server {
	addr := viper.GetString("server.addr")
	port := viper.GetString("server.port")

	if len(port) == 0 {
		port = "http"
	}

	redirectUrl, err := url.Parse(
		viper.GetString("openID.redirectUrl"),
	)

	if err != nil {
		logrus.WithError(err).Fatal("openID.redirectUrl is invalid")
	}

	router := mux.NewRouter()

	router.Methods("GET").Path("/steam/openid").Handler(
		server.NewOpenIdHandler(*redirectUrl),
	)

	return server.NewServer(
		addr,
		port,
		server.NewErrorHandler(router),
	)
}
