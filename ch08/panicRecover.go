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
	for _, val := range []int{1,2,0,4} {
		myDiv(val)
	}
}