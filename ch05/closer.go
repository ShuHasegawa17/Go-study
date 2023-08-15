package main

import (
	"fmt"
	"sort"
)

func main() {

	type Person struct {
		name string
		age  int
	}
	person := []Person{
		{"hase", 30},
		{"taro", 20},
		{"hana", 40},
	}
	// クロージャを引数にする
	// sort.Slice(ソート対象,ソート条件-クロージャ)）
	sort.Slice(person, func(i int, j int) bool {
		// iとjは比較対象のインデックス
		return person[i].age < person[j].age
	})
	fmt.Print(person) //年齢の昇順
}
