package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port string = ":8000"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})
	log.Println("Server Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
