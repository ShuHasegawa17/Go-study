package main

import "fmt"

func main() {
	val := 10
	mutableFunc(&val) // ポインタ（アドレス）を渡す
	fmt.Println("main",val) // 更新される→1000
}

func mutableFunc(pointer *int) {
	*pointer = 1000 // ポインタ（アドレス）参照先の値を変更
	fmt.Println("func",*pointer) // 1000

	// ちなみに以下の場合は変更されない
	//val := 1000
	//pointer = &val	// ポインタ（アドレス）を変更しているので元の値は変更されない。 
}