package utils

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestDBConnection(t *testing.T) {
	db := createConn()
	if db == nil {
		panic("")
	}
}

func TestGetData(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer fmt.Println(k)
	db := createConn()
	var wg sync.WaitGroup
	wg.Add(1)
	var wg2 sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg2.Add(1)
		go func(count int) {
			wg.Wait()
			defer wg2.Done()
			if db == nil {
				t.Error("Fail")
				//t.Failed()
				return
			}
			rows, err := db.Query("select * from user where id = " + strconv.Itoa(count))
			if err != nil {
				t.Error(err, ":", count)
			}
			for rows.Next() {
				var id int
				var name string
				rows.Scan(&id, &name)
				if id != 0 {
					inc()
				}
			}
			rows.Close()
		}(i)
	}
	wg.Done()
	wg2.Wait()
}

var k = 0
var a sync.Mutex

func inc() {
	a.Lock()
	k++
	defer a.Unlock()
}
