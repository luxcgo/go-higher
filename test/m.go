package main

import "os"

func main() {
	a := os.Stderr
	a.WriteString("hhh")
}
