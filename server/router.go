package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	apiRouter map[string]Handler
	port      string
)

type Handler struct {
	Path 		string
	Method      []string
	ServiceName string
}

func ServeHTTP(ApiRouter *map[string]Handler, portServe string) {
	apiRouter, port, *ApiRouter = *ApiRouter, portServe, nil
	r := mux.NewRouter()
	for path, info := range apiRouter {
		for _, val := range info.Method {
			r.HandleFunc(path, handle).Methods(val)
		}
	}
	log.Fatalln(http.ListenAndServe(":"+port, r))
}

func handle(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("123123123"))
}