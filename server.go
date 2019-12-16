package goweb

import (
	"fmt"
	"net/http"
	"log"
)

func CreateServer(port string, router string) {
	fmt.Printf("Starting server...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}