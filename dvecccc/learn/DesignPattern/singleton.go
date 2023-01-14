package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var initialized uint32

type Singleton struct {
}

func GetInstance1() *Singleton {
	mu := &sync.Mutex{}
	return getInstance(mu)
}

func getInstance(mu *sync.Mutex) *Singleton {
	var instance *Singleton
	if atomic.LoadUint32(&initialized) == 1 {
		fmt.Println("Singleton has initialized")
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if atomic.LoadUint32(&initialized) == 0 {
		instance = &Singleton{}
		atomic.StoreUint32(&initialized, 1)
		fmt.Println("Singleton is initialized first")
	}
	return instance
}

var once sync.Once

func GetInstance2() *Singleton {
	var instance *Singleton
	once.Do(func() {
		fmt.Println("first generation")
		instance = &Singleton{}
	})
	return instance
}

func main() {
	GetInstance1()
	GetInstance1()
	GetInstance1()
	GetInstance1()
	GetInstance1()
	GetInstance2()
	GetInstance2()
	GetInstance2()
}
