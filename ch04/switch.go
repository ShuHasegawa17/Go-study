package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.NewSource(time.Now().UnixNano())

	// ブロック内のみ有効な変数「n」
	switch n := rand.Intn(12); n {
	case 2,4,6,8,10:
		fmt.Println("10までの偶数です。",n)
	case 1,3,5,7,9:
		fmt.Println("10までの奇数です。",n)
	default:
		fmt.Println("その他。",n)
	}
	// ブロック外ではエラーとなる
	//fmt.Println("nは",n)

	// ブランクswitch
	// switch内で条件式を書かず、case内に書くこともできる
	switch n := rand.Intn(12); {
	case n % 2 == 0 && n <= 10:
		fmt.Println("10までの偶数です。",n)
	case n % 2 != 0 && n <= 10:
		fmt.Println("10までの奇数です。",n)
	default:
		fmt.Println("その他。",n)
	}


}