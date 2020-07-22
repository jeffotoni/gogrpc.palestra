package main

import (
	"log"
	"net/http"
)

// openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
// certbot certonly --standalone-supported-challenges http-01 -d your-domain.com
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/auth", Auth) // post
	mux.HandleFunc("/user", User) // put / post / get all
	mux.HandleFunc("/user/uuid", UserOne) // get one / delete

	server := &http.Server{
		Addr:    "0.0.0.0:443",
		Handler: mux,
	}
	log.Println("Serever Run:443")
	log.Fatal(server.ListenAndServeTLS("/etc/letsencrypt/live/http3.guula.com.br/fullchain.pem",
		"/etc/letsencrypt/live/http3.guula.com.br/privkey.pem"))
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
