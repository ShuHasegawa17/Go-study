package main

import "fmt"

func main() {
	x := "hello"
	ponterX := &x //ポインタ値の取得
	x2 := *ponterX // ポインタ値から値の取得（逆参照）
	fmt.Println(ponterX) //0x1400010a230（アドレス）
	fmt.Println(x2) // hello
	
	var pointer *int // ポインタ型の宣言。初期値nil
	fmt.Println(pointer) // nil
	//fmt.Println(*pointer) // パニック
}