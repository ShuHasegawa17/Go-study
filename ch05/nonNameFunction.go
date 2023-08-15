package main

import "fmt"

func main() {
	//関数は値であるため変数に代入できる
	var add = addFunc
	fmt.Println(add(10, 20))

	for i := 0; i < 5; i++ {
		//無名関数-即時実行()する
		func(j int) {
			fmt.Println("無名関数を実行", j)
		}(i)
	}

}

func addFunc(val1 int, val2 int) int {
	return val1 + val2
}
