package main

import "fmt"

func main() {
	fmt.Println(div(10, 2))
	fmt.Println(addTo(10, 1, 2, 3, 4, 5))

}

func div(a int, j int) int {
	return a / j
}

// 型はまとめて宣言できる
func div2(a, j int) int {
	return a / j
}

// ...可変長引数（スライス）も可能
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
