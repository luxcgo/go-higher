package main

import "log"

func main() {
	sl := make([]int, 10)
	log.Println(len(sl), cap(sl))
	ssl := sl[10:]
	log.Println(ssl)
}
