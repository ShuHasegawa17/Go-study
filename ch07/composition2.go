package main

import "fmt"

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	result := i.A * 2
	// ここで呼び出されるのは、innerのIntPrinter
	return i.IntPrinter((result))
}

type Outer struct {
	Inner // 埋め込み
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double()) //Inner: 20
}

