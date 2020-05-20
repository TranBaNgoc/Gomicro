package main

import (
	"net/http"
	"sync"
)

func main() {
	//cli := proto.NewHelloClient("hello", client.NewClient())
	//body, err := cli.Say(context.Background(), &proto.Request{})
	//if err == nil {
	//	fmt.Println(body.Message)
	//} else {
	//	fmt.Println(err)
	//}
	testHttp()
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
