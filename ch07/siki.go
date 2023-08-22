package main

import "fmt"

type Adder struct {
	value int
}
func (a Adder)AddTo(val int)int {
	return a.value + val
}

func main() {
	myAdder := Adder{value: 10}
	// メソッド値
	f1 := myAdder.AddTo
	fmt.Println((f1(20)))

	// メソッド式
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder,25))
}