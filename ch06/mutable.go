package main

import "fmt"

func main() {
	val := 10
	mutableFunc(&val)
	fmt.Println("main",val)
}

func mutableFunc(pointer *int) {
	*pointer = 1000
	fmt.Println("func",*pointer)
}