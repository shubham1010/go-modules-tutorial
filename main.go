package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/shubham1010/go-module-tutorial/routers"
)

func main() {
	log.Info("Starting...")

	router := routers.InitRoutes()
	log.Info("Started InitRoutes()...")

	server := &http.Server{
		Addr: 8080,
		Handler: router,
	}
	log.Info("Listening on port 8080...")
	server.ListenAndServe()
}
