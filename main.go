package main

import (
	"log"
	"net/http"

	"github.com/NatalyMac/test/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mix", api.GetMix).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
