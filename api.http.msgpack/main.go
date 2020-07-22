package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//go:generate msgp

type Login struct {
	Name       	string `msg:"name"`
	Lastname    string `msg:"lastname"`
	Age   	 	int `msg:"age"`
}

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
	if strings.ToUpper(r.Method) == http.MethodPost {
		b, err := ioutil.ReadAll(r.Body)
		if err !=nil {
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte(`{"msg":"error ReadAll"}`))
			return
		}

		var pp = Login{}
		err = json.Unmarshal(b, &pp)
		if err !=nil {
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte(`{"msg":"error json.Unmarshal"}`))
			return
		}

		fmt.Println("............. json recive.............")
		fmt.Println(pp)
		bts, err := pp.MarshalMsg(nil)
		if err !=nil {
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte(`{"msg":"error json.MarshalMsg"}`))
			return
		}

		var p = Login{}
		_, err = p.UnmarshalMsg(bts)
		if err !=nil {
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte(`{"msg":"error msgp.Unmarshaler"}`))
			return
		}
		println("...................done success MarshalMsg......................")
		fmt.Println(string(bts))
		println("...................done success Unmarshaler......................")
		fmt.Println(p.Name)
		fmt.Println(p.Lastname)
		fmt.Println(p.Age)

		w.WriteHeader(200)
		return
	}

	w.WriteHeader(400)
	w.Write([]byte(`{"msg":"user"}`))
}

func UserOne(w http.ResponseWriter, r *http.Request) {
	// ...
	w.WriteHeader(200)
	w.Write([]byte(`{"msg":"user one"}`))
}
