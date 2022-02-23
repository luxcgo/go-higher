package main

import (
	"fmt"
	"os"
	"sync/atomic"
)

var v atomic.Value

var f, _ = os.Create("out.txt")

type IceCreamMaker interface {
	// Great a customer.
	Hello()
}

type Jerry struct {
	// id   int
	name string
}

func (b *Ben) Hello() {
	s := fmt.Sprintf("Ben says, \"Hello my name is %s\"\n", *b.field1)
	f.WriteString(s)
}

type Ben struct {
	// name string
	field1 *[3]byte
	// field2 int
}

func (j *Jerry) Hello() {
	s := fmt.Sprintf("Jerry says, \"Hello my name is %s\"\n", j.name)
	f.WriteString(s)
}

func main() {
	slice := []byte("Ben")
	var arr [3]byte
	copy(arr[:], slice[:3])
	var jerry = &Jerry{
		"Jerry",
	}
	var ben = &Ben{field1: &arr}

	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}
