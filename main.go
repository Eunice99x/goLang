package main

import (
	"fmt"
	"html/template"
	"net/http"

	"younes.dev/go/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request){
	html, err := template.ParseFiles("templates/index.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	html.Execute(w, data.GetAll()[1])
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello",handleHello)
	server.HandleFunc("/template",handleTemplate)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8888", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}