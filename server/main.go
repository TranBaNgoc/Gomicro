package main

import (
	"github.com/micro/go-micro"
	h "gomicro/server/handler"
	proto "gomicro/server/pb"
	"gomicro/server/rest"
	"log"
)

func main() {
	service := micro.NewService(micro.Name("hello"))
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(h.Hello))
	rest.NewMatcher(APIMatcher)
	go rest.HTTPServe("8001")
	log.Fatal(service.Run())
}

var APIMatcher = map[string]rest.Handler{
	"/": {
		Request:     &proto.Request{},
		Response:    &proto.Response{},
		Handler:     new(h.Hello),
		ServiceName: "Say",
		Method:      "GET",
	},
	"/test/{id}": {
		Request:     &proto.Request{},
		Response:    &proto.Response{},
		Handler:     new(h.Hello),
		ServiceName: "Goodbye",
		Method:      "GET",
	},
}
