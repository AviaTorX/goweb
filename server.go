package goweb

import (
	"fmt"
	"net/http"
	"log"
)

func firstConnectioon(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Fprintf(w, "Hi Axi")
}

func CreateServer(port string, router string) {
	fmt.Printf("Starting server...\n")
	http.HandleFunc("/", firstConnectioon)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}