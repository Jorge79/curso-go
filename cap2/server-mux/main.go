package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux2 := http.NewServeMux()
	http.ListenAndServe(":8080", mux)

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})

	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Jorge!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
