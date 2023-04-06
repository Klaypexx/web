package main

import (
	"log"
	"net/http"
)

const (
	port = ":2020"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index)     //если url будет /home, то он вызывает функцию index
	mux.HandleFunc("/home/post", post) //если url будет /post, то он вызывает функцию post

	// Реализуем отдачу статики
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Start server " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
