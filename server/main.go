package main

import (
	"github.com/micro/go-micro/v2"
	"gomicro/server/business"
	proto "gomicro/server/protobuf"
	"log"
)

func main() {
	service := micro.NewService(micro.Name("noti"))
	service.Init()
	proto.RegisterUserServiceHandler(service.Server(), new(business.UserService))

	log.Fatal(service.Run())
}
