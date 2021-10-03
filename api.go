package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetRqstHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	message := fmt.Sprintf("App Id: %s", os.Getenv("APP_ID"))

	// w.Write([]byte(message))
	json.NewEncoder(w).Encode(message)
}

func PostRqstHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	message := fmt.Sprintf("App Id: %s", os.Getenv("APP_ID"))

	// w.Write([]byte(message))
	json.NewEncoder(w).Encode(message)
}

func main() {
	// load env values
	godotenv.Load(".env")

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/get", GetRqstHandler).Methods("GET")
	r.HandleFunc("/post", PostRqstHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
