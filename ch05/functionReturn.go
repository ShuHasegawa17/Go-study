package main

import "fmt"

// 複数の戻り値
func multiAddFunc(add1 int, add2 int) (int, string) {
	result := add1 + add2
	return result, "足し算の結果"
}

// 名前付き戻り値
func funcName() (message string) {
	message = "名前付き戻り値"
	// 別の値を返しても良い
	return "別の値を返す"
}

func brankFunc() (message string) {
	message = "ブランクreturn"
	return
}

func main() {
	reslt, message := multiAddFunc(10, 20)
	fmt.Println(reslt, message)

	fmt.Println(funcName())

	fmt.Println(brankFunc())
}
