package rest

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
	port int
)


type Handler struct {
	Request     interface{}
	Response    interface{}
	Handler     interface{}
	Method      string
	ServiceName string
}

func HTTPServe(port string) {
	r := mux.NewRouter()
	for path := range apiRouter {
		r.HandleFunc(path, process).Methods(apiRouter[path].Method)
	}
	log.Fatalln(http.ListenAndServe(":" + port, r))
}

func NewMatcher(mapApi map[string]Handler) {
	apiRouter = mapApi
}

func getRouter(path string) Handler {
	if res, ok := apiRouter[path]; ok == true {
		return res
	}

	return Handler{}
}

func process(writer http.ResponseWriter, request *http.Request) {
		
	match := getRouter(request.RequestURI)
	req := match.Request
	res := match.Response
	err := json.NewDecoder(request.Body).Decode(&req)
	defer request.Body.Close()
	params := []interface{}{req, res}
	err = call(match.Handler, params, match.ServiceName)
	if err != nil {
		log.Println("Call function:", err)
	}
	result, _ := json.Marshal(res)
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(result)
	if err != nil {
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
		object := objects[t]
		in[i] = reflect.ValueOf(object)
	}
	method.Call(in)
	return nil
}

