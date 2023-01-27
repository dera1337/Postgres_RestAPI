package main

import (
	"encoding/json"
	"exercisee/db"
	"fmt"
	"io"
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

	r.HandleFunc("/create/class", insertClassHandler).Methods("POST")

	r.HandleFunc("/create/class", updateClassHandler).Methods("UPDATE")

	r.HandleFunc("/create/class", deleteClassHandler).Methods("DELETE")

	r.HandleFunc("/create/class", readClassHandler).Methods("DELETE")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}

}

type GenericResponse struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

func writeResponse(w http.ResponseWriter, data interface{}, msg string, statusCode int) {
	response := GenericResponse{
		Data:       data,
		Message:    msg,
		StatusCode: statusCode,
	}

	jsonBytes, _ := json.Marshal(&response)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}

func readClassHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var class db.Class

	err = json.Unmarshal(jsonBytes, &class)
	if err != nil {
		return
	}

	class, err = db.ReadRow(db.Conn, 1)
	if err != nil {
		return
	}
}

func updateClassHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var class db.Class

	err = json.Unmarshal(jsonBytes, &class)
	if err != nil {
		return
	}

	err = db.UpdateRow(db.Conn, class)
	if err != nil {
		return
	}
}

func deleteClassHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var class db.Class

	err = json.Unmarshal(jsonBytes, &class)
	if err != nil {
		return
	}

	err = db.DeleteRow(db.Conn, class)
	if err != nil {
		return
	}
}

func insertClassHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var class db.Class

	err = json.Unmarshal(jsonBytes, &class)
	if err != nil {
		return
	}

	err = db.InsertRow(db.Conn, class)
	if err != nil {
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
