package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/auth", Auth) // post
	mux.HandleFunc("/user", User) // put / post / get all
	mux.HandleFunc("/user/uuid", UserOne) // get one / delete

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Println("Serever Run:8080")
	log.Fatal(server.ListenAndServe())
}

func Auth(w http.ResponseWriter, r *http.Request) {
	// ...
	w.WriteHeader(200)
	w.Write([]byte(`{"msg":"auth"}`))
}

func User(w http.ResponseWriter, r *http.Request) {
	// ...
	w.WriteHeader(200)
	w.Write([]byte(`{"msg":"user"}`))
}

func UserOne(w http.ResponseWriter, r *http.Request) {
	// ...
	w.WriteHeader(200)
	w.Write([]byte(`{"msg":"user one"}`))
}
