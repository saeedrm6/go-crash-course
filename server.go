package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang-rest-api/routes"
	"log"
	"net/http"
)

const port string = ":8000"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts/add", routes.AddPost).Methods("POST")
	log.Println("Server Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}