package main

import "fmt"

func main() {
	// 配列の作成
	var l1 = [3]int{10,20,30}
	var l2 = [...]string{"aa","bb"}
	// スライス
	var s1 = []int{10,20,30}
	// スライスは追加できる（戻り値が必要）
	s1 = append(s1, 40);

	// スライスのコピー
	s := []int{10,20,30,40}
	// makeでコピー元と同じ長さの配列を用意
	s2 := make([]int, len(s))
	// copyを実施
	copy(s2,s)
	s2 = append(s2, 50)

	fmt.Println(l1,l2,s1)
	fmt.Println(s, s2)
}