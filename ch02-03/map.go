package main

import "fmt"

func main() {

	// 宣言
	var m1 map[int]string  // NG: nilマップ。値の追加不可
	m2 := map[int]string{} // OK: 空マップの追加。値の追加可能

	// 値の読み書き
	m2[1] = "値1"
	m2[2] = "値2"
	m2[3] = "値3"
	fmt.Println(m2[2]) // 値2

	// 削除
	delete(m2, 2)
	fmt.Println(m2[2])   // 値なし：nil
	fmt.Println(len(m2)) // 2

	fmt.Println(m1, m2)

	// カンマOKイディオム
	m3 := map[string]int{
		"first": 0,
		"second": 1,
	}
	val, ok := m3["first"]
	fmt.Println(val,ok) // 0, true → キーが存在し、値が0
	val2, ok2 := m3["notkey"]
	fmt.Println(val2,ok2) // 0, false　→　キーがなく、値はデフォルト0
	
	

}
