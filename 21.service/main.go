package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	var s string = r.URL.String()
	u, err := r.URL.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	len := len(u.Scheme) + len(u.Host)
	fmt.Println(len, u.Host, u.Scheme)
	fmt.Fprintf(w, "Title: %s", u.Path) // send data to client side
}

func main() {
	http.HandleFunc("/books/", sayhelloName) // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
