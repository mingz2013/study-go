package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello!\n")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:8000", h)
	if err != nil {
		log.Fatal(err)
	}
}
