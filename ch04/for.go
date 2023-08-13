package main

import "fmt"

func main() {
	// 通常のfor文
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 拡張for文
	evenVal := []int{1,2,3,4,5,6}
	// index不要の場合「_」にする。なお、2つめ（value）が不要の場合は宣言しないでよい。
	for _, v := range evenVal {
		if v % 2 == 0 {
			fmt.Println(v)
		}
	}

	// 条件のみfor文（while文）
	i := 2
	for i < 10 {
		fmt.Println(i)
		i = i * 2
	}

	// 条件なしfor文（無限ループ）
	n := 1;
	for {
		fmt.Println(n)
		if n > 5 {
			break
		}
		n++
	}
}
