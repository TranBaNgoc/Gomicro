package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
)

var (
	apiRouter map[string]Handler
	port      string
)

type Handler struct {
	Request     interface{}
	Response    interface{}
	Handler     interface{}
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

func getRouter(request *http.Request) (Handler, error) {
	if handle, ok := apiRouter[request.RequestURI]; ok == true {
		return handle, nil
	}
	if handle, err := mux.CurrentRoute(request).GetPathTemplate(); err == nil {
		return apiRouter[handle], nil
	} else {
		return Handler{}, err
	}
}

func handle(writer http.ResponseWriter, request *http.Request) {
	match, err := getRouter(request)
	if err != nil {
		log.Println(55, err)
		return
	}

	req, res := match.Request, match.Response
	if request.Method == "GET" {
		vars := mux.Vars(request)
		paramsJson, err := json.Marshal(vars)
		if err != nil {
			log.Fatalln(64, err)
		}
		err = json.Unmarshal(paramsJson, &req)
		if err != nil {
			log.Fatalln(71, err)
		}
	} else {
		decoder := json.NewDecoder(request.Body)
		err = decoder.Decode(&req)
		if err != nil {
			log.Fatalln(77, err)
		}
		request.Body.Close()
	}

	functionInput := []interface{}{req, res}
	if call(match.Handler, functionInput, match.ServiceName) != nil {
		log.Println("Call function:", err)
	}

	result, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(86, err)
	}

	writer.Header().Set("Content-Type", "application/json")
	if _, err = writer.Write(result); err != nil {
		log.Fatalln("Response:", err)
	}
}

func call(handler interface{}, input []interface{}, methodName string) error {
	objects := make(map[reflect.Type]interface{})
	for _, val := range input {
		objects[reflect.TypeOf(val)] = val
	}

	method := reflect.ValueOf(handler).MethodByName(methodName)

	in := make([]reflect.Value, method.Type().NumIn())
	in[0] = reflect.ValueOf(context.Background())
	for i := 1; i < method.Type().NumIn(); i++ {
		t := method.Type().In(i)
		in[i] = reflect.ValueOf(objects[t])
	}

	if err := method.Call(in); err[0].Interface() != nil {
		return err[0].Interface().(error)
	} else {
		return nil
	}
}
