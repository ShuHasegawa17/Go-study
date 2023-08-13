package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	// ブロック内のみ有効な変数「n」
	if n := rand.Intn(10); n == 5 {
		fmt.Println("nは5", n)
	} else if n < 5 {
		fmt.Println("nは5未満", n)
	} else {
		fmt.Println("nは5より上", n)
	}
	// ブロック外ではエラーとなる
	//fmt.Println("nは",n)
}
