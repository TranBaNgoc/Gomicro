package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	pb "gomicro/client/protobuf"
	"net/http"
	"sync"
)

func main() {
	cli := pb.NewUserServiceClient("user.service", client.NewClient())
	body, err := cli.GetListUsers(context.Background(), &pb.EmptyRequest{})
	if err == nil {
		fmt.Println(body.User)
	} else {
		fmt.Println(err)
	}

	//body2, err2 := cli.AddUser(context.Background(), &protobuf.User{
	//	Name:                 "Nguyen ba Nam",
	//	Age:                  22,
	//})
	//if err2 == nil {
	//	fmt.Println(body2.Message)
	//} else {
	//	fmt.Println(err)
	//}
	//testHttp()
}

func testHttp() {
	var wg sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			_, err := http.Get("http://localhost:8001/test/1")
			if err != nil {
				panic(err)
			} else {
				//fmt.Println(resp)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
