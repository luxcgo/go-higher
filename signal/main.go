package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"
)

func count() {
	println(runtime.NumGoroutine())
}

func main() {
	count()
	defer count()
	// run program
	go run()

	// allow graceful termination
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// cleanup if needed here somewhere
	fmt.Println("Have a nice day!")

	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

}

func run() {
	for {
		log.Println(time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}
