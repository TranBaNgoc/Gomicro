package test

import (
	"gomicro/server/model"
	"gotest.tools/assert"
	"sync"
	"testing"
)

var (
	successCounter = 0
	m              sync.Mutex
)

func inc() {
	m.Lock()
	defer m.Unlock()
	successCounter++
}

func TestQuery(t *testing.T) {
	var (
		done sync.WaitGroup
		wg   sync.WaitGroup
	)

	n := 1000
	wg.Add(1)
	for i := 0; i < n; i++ {
		done.Add(1)
		go func(age int64) {
			wg.Wait()
			defer done.Done()
			newUser := &model.User{
				Name: "Tran Ba Ngoc",
				Age:  age,
			}
			if err := model.AddUser(newUser); err == nil {
				inc()
			}
		}(int64(i))
	}
	wg.Done()
	done.Wait()
	assert.Equal(t, successCounter, n, "%d queries success", successCounter)
}
