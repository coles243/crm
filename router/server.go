package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {

	r := mux.NewRouter()

	r.HandleFunc("/", GetAll)

	http.ListenAndServe(":3000", r)

}
