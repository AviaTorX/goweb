package goweb

import (
	"fmt"
	"net/http"
	"log"
)

func firstConnectioon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Axi")
}

func CreateServer(port string, router string) {
	fmt.Printf("Starting server...\n")
	http.HandleFunc("/", firstConnectioon)
	log.Fatal(http.ListenAndServe(":8080", nil))
}