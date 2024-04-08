package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/sarulabs/di"
	"go.dev.io/api"
	"go.dev.io/data"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello, HTTP!\n")
}

func getTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		fmt.Printf("Template error: %v", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = html.Execute(w, data.GetAll())
	if err != nil {
		return
	}
}


func main() {
    app := createApp()
    defer app.Delete()

    server := app.Get("http-server").(*http.Server)
    err := server.ListenAndServe()
    if err != nil {
        fmt.Println("Error opening the server:", err)
    }
}

func createApp() di.Container {
    builder, _ := di.NewBuilder()

    builder.Add([]di.Def{
        {
            Name:  "http-server",
            Build: func(ctn di.Container) (interface{}, error) {
                mux := http.NewServeMux()
                mux.HandleFunc("/hello", ctn.Get("hello-handler").(http.HandlerFunc))
                mux.HandleFunc("/templates", ctn.Get("template-handler").(http.HandlerFunc))
                mux.HandleFunc("/api/exhibitions", ctn.Get("api-get-handler").(http.HandlerFunc))
                mux.HandleFunc("/api/exhibitions/new", ctn.Get("api-post-handler").(http.HandlerFunc))
                fs := http.FileServer(http.Dir("./public"))
                mux.Handle("/", fs)
                return &http.Server{
                    Addr:    ":8888",
                    Handler: mux,
                }, nil
            },
            Close: func(obj interface{}) error {
                return obj.(*http.Server).Close()
            },
        },
        {
            Name:  "hello-handler",
            Build: func(ctn di.Container) (interface{}, error) {
                return http.HandlerFunc(getHello), nil
            },
        },
        {
            Name:  "template-handler",
            Build: func(ctn di.Container) (interface{}, error) {
                return http.HandlerFunc(getTemplate), nil
            },
        },
        {
            Name:  "api-get-handler",
            Build: func(ctn di.Container) (interface{}, error) {
                return http.HandlerFunc(api.Get), nil
            },
        },
        {
            Name:  "api-post-handler",
            Build: func(ctn di.Container) (interface{}, error) {
                return http.HandlerFunc(api.Post), nil
            },
        },
        {
            Name:  "exhibition-data",
            Build: func(ctn di.Container) (interface{}, error) {
                return data.GetAll(), nil
            },
        },
    }...)

    return builder.Build()
}
