package controllers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is login page"))
}
