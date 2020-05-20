package main

import (
	"github.com/micro/go-micro"
	business "gomicro/server/business"
	proto "gomicro/server/pb"
	"log"
)

var APIMatcher = map[string]Handler{
	"/hello": {
		Request:     &proto.Request{},
		Response:    &proto.Response{},
		Handler:     &business.Hello{},
		ServiceName: "Say",
		Method:      []string{"GET"},
	},
	"/goodbye/{Name}": {
		Request:     &proto.Request{},
		Response:    &proto.Response{},
		Handler:     &business.Hello{},
		ServiceName: "Goodbye",
		Method:      []string{"GET", "POST"},
	},
}

func main() {
	go ServeHTTP(&APIMatcher, "8001")
	service := micro.NewService(micro.Name("hello"))
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(business.Hello))
	log.Fatal(service.Run())
}
