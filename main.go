package main

import (
	"encoding/json"
	"exercisee/db"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.CloseConnection()

	r := mux.NewRouter()

	r.HandleFunc("/create/class", insertClassHandler).Methods("POST")

	// r.HandleFunc("/create/class", updateClassHandler).Methods("UPDATE")

	r.HandleFunc("/delete/class/{classID}", deleteClassHandler).Methods("DELETE")

	// r.HandleFunc("/create/class", readClassHandler).Methods("GET")

	err = http.ListenAndServe("localhost:8080", r)
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
	params := mux.Vars(r)
	idString := params["classID"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		writeResponse(w, nil, "expecting int for classID, but got string", http.StatusBadRequest)
		return
	}

	err = db.DeleteRow(db.Conn, id)
	if err != nil {
		messageDelete := fmt.Sprintf("can't delete row with ID: %d", id)
		writeResponse(w, nil, messageDelete, http.StatusBadRequest)
		return
	}

	messageOk := fmt.Sprintf("deleted row with ID: %d", id)
	writeResponse(w, nil, messageOk, http.StatusOK)
}

func insertClassHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, nil, "expecting content, please type something", http.StatusBadRequest)
		return
	}

	var class db.Class

	err = json.Unmarshal(jsonBytes, &class)
	if err != nil {
		writeResponse(w, nil, "invalidJSON", http.StatusBadRequest)
		return
	}

	err = db.InsertRow(db.Conn, class)
	if err != nil {
		messageFailInsert := fmt.Sprintf("Fail inserting: %s", class.Name)
		writeResponse(w, nil, messageFailInsert, http.StatusBadRequest)
		return
	}
	messageSuccessInsert := fmt.Sprintf("Success inserting class with name: %s", class.Name)
	writeResponse(w, nil, messageSuccessInsert, http.StatusOK)
}

func handleParameters(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	fmt.Println(pathParams)

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	fmt.Printf("limit: %s\n", limit)
	fmt.Printf("offset: %s\n", offset)
}
