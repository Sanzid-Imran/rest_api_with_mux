package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")

	message := fmt.Sprintf("App Id: %s", os.Getenv("APP_ID"))
	w.Write([]byte(message))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", GetHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
