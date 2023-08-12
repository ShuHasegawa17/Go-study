package main

import "fmt"

func main() {
	// 整数
	var i1 int = 10
	var i2 = 10
	// 浮動小数
	var f1 float64 = 10.5
	var f2 = 10.5
	// bool
	var b1 bool = true
	var b2 = false
	// string
	var s1 = "hello"
	var s2 string = "hello2"
	// rune
	var r1 = 'h'
	var r2 rune = 'h'

	// 型のキャスト
	//f1 = i1　//暗黙的キャストはNG
	f1 = float64(i1)
	b1 = s1 == s2

	// NG
	castNum := 10
	//castStr := string(castNum)

	// 一括宣言
	var i3, s3 = 20, "abc"
	var (
		i4     = 10
		i5     int
		s4, s5 = "aa", "bb"
	)

	// 関数内での変数宣言
	i6 := 10

	// 定数
	const con1 = "定数"
	const (
		con2 = "const"
		con3 = "const2"
	)

	fmt.Println(i1, i2, f1, f2)
	fmt.Println(b1, b2, s1, s2, r1, r2)
	fmt.Println(i3, s3, i4, i5, s4, s5, i6)
	fmt.Println(castNum)

}
