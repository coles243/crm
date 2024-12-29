package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {

	r := mux.NewRouter()
	r.HandleFunc("/customers", GetAll)
	r.HandleFunc("/userid/{id}/", GetCustomer)
	r.HandleFunc("/customer", CreateCustomer).Methods("POST")
	r.HandleFunc("/delete/customer/{id}/", RemoveCustomer).Methods("DELETE")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalln(err)
	}

}
