package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Format("[2006-01-02 15-04-05]"))

	// t := jitterbug.New(
	// 	time.Second*4,
	// 	&jitterbug.Norm{Stdev: 0},
	// )

	// // jitterbug.Ticker behaves like time.Ticker
	// for tick := range t.C {
	// 	log.Println(tick)
	// }
}
