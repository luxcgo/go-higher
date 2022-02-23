package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func listenSignal() {
	defer wg.Done()
	// log.Printf("收到结束信号\n")
	// time.Sleep(time.Second * 3)
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	for sig := range ch {
		log.Printf("收到结束信号(%s)，准备结束进程\n", sig.String())
		time.Sleep(time.Second * 3)
		return
	}
}

func count() {
	println(runtime.NumGoroutine())
}

func main() {
	count()
	defer count()
	wg.Add(1)
	go listenSignal()
	wg.Wait()
}
