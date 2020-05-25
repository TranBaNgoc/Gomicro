package main

import (
	"github.com/micro/go-micro"
	"gomicro/server/business"
	proto "gomicro/server/protobuf"
	"log"
)

var APIMatcher = map[string]Handler{
	"/all": {
		Request:     &proto.EmptyRequest{},
		Response:    &proto.GetListUserResponse{},
		Handler:     &business.UserService{},
		ServiceName: "GetListUsers",
		Method:      []string{"GET"},
	},
	"/add": {
		Request:     &proto.User{},
		Response:    &proto.AddUserResponse{},
		Handler:     &business.UserService{},
		ServiceName: "AddUser",
		Method:      []string{"POST"},
	},
}

func main() {
	go ServeHTTP(&APIMatcher, "8001")
	service := micro.NewService(micro.Name("user.service"))
	service.Init()
	proto.RegisterUserServiceHandler(service.Server(), new(business.UserService))
	log.Fatal(service.Run())
}
