package main

import (
	"log"
	"net/http"
	"Server/service"
)

func main() {
	log.Println("Server started")
	server := service.NewRouter() // get server
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", server))
}