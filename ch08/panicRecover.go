package main

import "fmt"

func myDiv(i int) {
	// defer内でリカバー処理を記述する
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("処理継続", v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 4} {
		// 0除算でパニックとなるがリカバーするため、処理が継続される。
		//60
		//30
		//処理継続 runtime error: integer divide by zero
		//15
		myDiv(val)
	}
}
