package main

import (
	"context"
	"fmt"
	pb "gomicro/client/protobuf"
	"net/http"
	"sync"
)

func main() {

	// create the greeter client using the service name and client
	greeter := pb.NewUserService()
	// request the Hello method on the Greeter handler
	rsp, err := greeter.GetListUsers(context.TODO(), &pb.EmptyRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.User)
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
