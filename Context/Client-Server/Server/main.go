package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	println("Requisição inicializada")
	defer log.Println("Requisição finalizada")

	select {
	case <-time.After(time.Second * 5):
		//Imprime no stdout
		log.Println("Requisição processada com sucesso")

		//Imprime no browser
		w.Write([]byte("Requisição processada com sucesso!"))
	case <-ctx.Done():
		//Imprime no stdout
		log.Println("Requisição cancelada pelo cliente")
	}
}
