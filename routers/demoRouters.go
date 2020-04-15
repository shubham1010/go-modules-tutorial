package routers

import (
	"github.com/gorilla/mux"
	"github.com/shubham1010/go-module-tutorial/controllers"
)

func SetRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/home", controllers.Home).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("GET")

	return router
}
