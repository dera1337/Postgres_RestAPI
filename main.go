package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// err := db.InitDB()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// defer db.CloseConnection()

	r := mux.NewRouter()
	r.HandleFunc("/{custID}", handleParameters)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func handleParameters(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	fmt.Println(pathParams)

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	fmt.Printf("limit: %s\n", limit)
	fmt.Printf("offset: %s\n", offset)
}
