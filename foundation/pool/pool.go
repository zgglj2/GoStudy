package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 5
)

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	fmt.Println("close connection")
	return nil
}

var idCounter int32

func createConnection() interface{} {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{ID: id}
}

func dbQuery(query int, pool *sync.Pool) {
	conn := pool.Get().(*dbConnection)
	defer pool.Put(conn)

	time.Sleep(time.Second)
	fmt.Printf("query %d, ID: %d\n", query, conn.ID)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p := &sync.Pool{
		New: createConnection,
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
}
