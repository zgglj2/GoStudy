package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type ReentrantLock struct {
	sync.Mutex
	recursion int32
	owner     int64
}

func GetGoroutineId() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func NewReentrantLock() sync.Locker {
	return &ReentrantLock{
		Mutex:     sync.Mutex{},
		recursion: 0,
		owner:     0,
	}
}

func (l *ReentrantLock) Lock() {
	gid := GetGoroutineId()
	if atomic.LoadInt64(&l.owner) == gid {
		l.recursion++
		return
	}
	l.Mutex.Lock()
	atomic.StoreInt64(&l.owner, gid)
	l.recursion = 1
}

func (l *ReentrantLock) Unlock() {
	gid := GetGoroutineId()
	if atomic.LoadInt64(&l.owner) != gid {
		panic("unlock a lock that is not locked by current goroutine")
	}

	l.recursion--
	if l.recursion != 0 {
		return
	}
	atomic.StoreInt64(&l.owner, 0)
	l.Mutex.Unlock()
}

func main() {
	var lock = &ReentrantLock{}
	lock.Lock()
	fmt.Println("Lock1")
	lock.Lock()
	fmt.Println("Lock2")
	lock.Unlock()
	fmt.Println("Unlock1")
	lock.Unlock()
	fmt.Println("Unlock2")

}
