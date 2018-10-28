package test_140_mux

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello home"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello home"))
}
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello home"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
