package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

type Ping struct {
	Status	string	`json:"status"`
	Message string `json:"message"`
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := json.NewEncoder(w).Encode(map[string]string{"200": "OK"})
	_, err := json.Marshal(res)

	if err != nil {
		res = json.NewEncoder(w).Encode(map[string]string{"500": "Internal Server Error."})
	}

	return
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexRoute).Methods("GET").Schemes("http", "https")
	http.Handle("/", router)

	s := &http.Server{
		Addr:	":8080",
	}

	fmt.Println("Initializing server...")

	log.Fatal(s.ListenAndServe())
}
