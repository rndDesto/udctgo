package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = []Customer{
	{1, "Badu", "Admin", "BaduKrenz@gmail.com", "081377263344", false},
	{2, "Ani", "Client", "ani99@gmail.com", "081366228844", true},
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	customerID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid customer ID"})
		return
	}

	for _, customer := range customers {
		if customer.ID == customerID {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	newCustomer.ID = len(customers) + 1
	customers = append(customers, newCustomer)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	customerID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid customer ID"})
		return
	}

	var updatedCustomer Customer
	err = json.NewDecoder(r.Body).Decode(&updatedCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	for i, customer := range customers {
		if customer.ID == customerID {
			customers[i] = updatedCustomer
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	customerID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid customer ID"})
		return
	}

	for i, customer := range customers {
		if customer.ID == customerID {
			customers = append(customers[:i], customers[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully"})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customers", getAllCustomers).Methods("GET")
	r.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", addCustomer).Methods("POST")
	r.HandleFunc("/customers/{id:[0-9]+}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id:[0-9]+}", deleteCustomer).Methods("DELETE")

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
