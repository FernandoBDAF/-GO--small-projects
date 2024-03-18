package main

import (
	"fmt"
	"net/http"
	"museum/handlers"
	"museum/api"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handlers.HandleHello)
	server.HandleFunc("/template", handlers.HandleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)

	staticServer := http.FileServer(http.Dir("./public"))
	server.Handle("/", staticServer)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}