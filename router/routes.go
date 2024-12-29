package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	var customer []Customer
	data, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(w, "Unable to initalize with Database")
		w.WriteHeader(http.StatusUnauthorized)

	}
	database := fmt.Sprintf("%v/router/database.json", data)
	file, err := os.ReadFile(database)
	if err != nil {
		fmt.Fprintln(w, "Unable to read from database")
		w.WriteHeader(http.StatusUnauthorized)
	}
	// create a list of structs to deseralize the json file
	json.Unmarshal(file, &customer)
	// then we will  reseralize the data as json content
	json.NewEncoder(w).Encode(customer)

}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	var customer []Customer
	// grab response from user
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	response := mux.Vars(r)

	// grab data from database
	datas, _ := os.Getwd()
	database := fmt.Sprintf("%v/router/database.json", datas)
	file, err := os.ReadFile(database)
	if err != nil {
		fmt.Fprintln(w, "Unable to read from database")
		w.WriteHeader(http.StatusUnauthorized)

	}

	intconversion, _ := strconv.Atoi(response["id"])

	// create a list of structs to deseralize the json file
	json.Unmarshal(file, &customer)

	for _, data := range customer {
		if data.ID == intconversion {
			json.NewEncoder(w).Encode(data)
		}
	}

	w.WriteHeader(http.StatusNotFound)

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var listofcustomers []Customer
	var customer Customer
	w.Header().Set("Content-Type", "application/json")

	// decode the body of the request
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Fprintln(w, "Unable to decoded")
	}

	datas, _ := os.Getwd()
	database := fmt.Sprintf("%v/router/database.json", datas)
	file, err := os.ReadFile(database)
	if err != nil {
		fmt.Fprintln(w, "Unable to read from database")
		w.WriteHeader(http.StatusUnauthorized)

	}

	// decode the list of users
	json.Unmarshal(file, &listofcustomers)

	for _, data := range listofcustomers {
		if data.Email == customer.Email {
			http.Error(w, "Duplicate email record", http.StatusForbidden)
			return
		}
	}

	// add unqiue ID
	customer.ID = len(listofcustomers) + 1
	// add new user
	listofcustomers = append(listofcustomers, customer)
	writebacktofile, _ := json.MarshalIndent(listofcustomers, "", "  ")
	write, _ := os.OpenFile(database, os.O_RDWR, 0644)

	defer write.Close()
	write.Write(writebacktofile)
	w.WriteHeader(http.StatusCreated)
}

func RemoveCustomer(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	w.Header().Set("Content-Type", "application/json")
	// grab data from the requests
	id := mux.Vars(r)

	data, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(w, "Unable to initalize with Database")
		w.WriteHeader(http.StatusUnauthorized)

	}

	idn, _ := strconv.Atoi(id["id"])

	database := fmt.Sprintf("%v/router/database.json", data)
	file, err := os.ReadFile(database)

	if err != nil {
		http.Error(w, "Unable to read in database", http.StatusRequestTimeout)
		return
	}

	json.Unmarshal(file, &customers)

	for i, data := range customers {
		if data.ID == idn {
			customers = slices.Delete(customers, i, i+1)
			break
		}
	}

	// json.NewEncoder(w).Encode(customers)

	writebacktofile, _ := json.MarshalIndent(customers, "", " ")
	write, _ := os.OpenFile(database, os.O_RDWR|os.O_TRUNC, 0644)

	defer write.Close()
	write.Write(writebacktofile)

	w.WriteHeader(204)
}
