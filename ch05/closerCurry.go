package main

import "fmt"

func main() {

	// カリー化
	twoBase := makeMult(2) //2倍する関数
	threeBase := makeMult(3) // 3倍する関数

	for i := 0; i<5; i++ {
		fmt.Println(i, "を2倍", twoBase(i))
		fmt.Println(i, "を3倍", threeBase(i))
 }

}

// クロージャを戻り値にする
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
