package main

import (
	"01.kood.tech/suzoagba/ascii-art-web/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ascii-art", handlers.FormHandler)
	fmt.Printf("Starting server at post: 8080\nhttp://localhost:8080/\n")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// if err != nil {
	// log.Fatal(err)
}