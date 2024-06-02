package main

import (
	"sync"
	"time"
)

type Singleton struct {
	Name string
}

var (
	instance *Singleton
	lock     sync.Mutex
	once     sync.Once
	onceInit sync.Once
)

func GetInstance() *Singleton {
	if instance == nil { // 19 行
		lock.Lock()          // 20 行
		defer lock.Unlock()  // 21 行
		if instance == nil { // 22 行
			instance = &Singleton{Name: "zhangsan"}
		}
	}
	return instance
}

func GetInstanceByOnce() *Singleton {
	once.Do(func() {
		onceInit.Do(func() {
			instance = &Singleton{Name: "once"}
		})
	})

	return instance
}

func main() {
	for i := 0; i < 500; i++ {
		go func() {
			_ = GetInstance()
			//_ = GetInstanceByOnce()
		}()
	}
	time.Sleep(time.Second)
}
