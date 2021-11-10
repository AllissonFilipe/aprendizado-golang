package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo !"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pagina Raiz"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar pagina de usuarios"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
