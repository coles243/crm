package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	var customer []Customer
	file, err := os.OpenFile("database.json", os.O_RDWR, 0664)
	// file, err := os.ReadFile("database.json")
	if err != nil {
		fmt.Fprintln(w, "Unable to read from database")
		w.WriteHeader(http.StatusUnauthorized)
	}

	defer file.Close()

	b1 := make([]byte, 1000000000)

	test, _ := file.Read(b1)
	// create a list of structs to deseralize the json file
	json.Unmarshal(string(test), &customer)

	// then we will  reseralize the data as json content
	json.NewEncoder(w).Encode(customer)

}
