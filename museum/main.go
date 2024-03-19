package main

import (
	"fmt"
	"museum/api"
	"museum/handlers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run()
	}()
	// if this is running out of the go routine it will block the main function here and the other server will not start

	server := http.NewServeMux()
	server.HandleFunc("/hello", handlers.HandleHello)
	server.HandleFunc("/template", handlers.HandleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	staticServer := http.FileServer(http.Dir("./public"))
	server.Handle("/", staticServer)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}