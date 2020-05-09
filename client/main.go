package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	proto "gomicro/server/pb"
)

func main() {
	cli := proto.NewHelloClient("hello", client.NewClient())
	body, err := cli.Say(context.Background(), &proto.Request{})
	if err == nil {
		fmt.Println(body.Message)
	} else {
		fmt.Println(err)
	}
}
