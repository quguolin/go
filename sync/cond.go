package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() //获取锁
	fmt.Println("block......", x)
	cond.Wait() //等待通知  暂时阻塞
	fmt.Println("unblock......", x)
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	cond.L.Unlock() //释放锁
}

func main() {
	for i := 0; i < 3; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 2)
	fmt.Println("Signal")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	fmt.Println("Broadcast")
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 2)
}
