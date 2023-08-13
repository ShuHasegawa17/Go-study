package main

import "fmt"

func main() {
	x := 10
	{
		fmt.Println(x) //10
		x := 5
		fmt.Println(x) //5 -> x:=10のシャドーイング発生
	}
	fmt.Println(x) // 10 → x:=10のシャドーイング解除
}
