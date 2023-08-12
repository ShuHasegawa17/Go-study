package main

import "fmt"

func main() {
	// 構造体の作成
	type human struct {
		name string
		age int
	}
	
	// 初期化
	var taro human
	hanako := human{} //Mapとは違い、varと同様に、0値で初期化。
	takashi := human{"たかし",20}
	yuki := human{name: "ゆき"}

	fmt.Println(taro) // { 0}
	fmt.Println(hanako) // { 0} 
	fmt.Println(takashi) // {たかし 20}
	fmt.Println(yuki) // {ゆき 0}

	// 値の設定
	hanako.name = "はなこ"
	fmt.Println(hanako) // { 0}

	// 無名構造体
	var game struct {
		title string
		price int
	}
	game.title = "FF"
	game.price = 7000

	// 無名構造体 - 初期化時に設定（リテラル代入）
	game2 := struct {
		title string
		memo string
	} {
		title: "DQ",
		memo: "楽しい",
	}
	fmt.Println(game2)

	type human1 struct {
		name string
		age int
	}
	h1 := human1{}

	type human2 struct {
		name string
		age int
	}
	h2 := human2{}

	type human3 struct {
		age int
		name string
	}
	h3 := human3{}

	// 構造体の構成が一致する場合、型変換で代入できる
	h1 = human1(h2) // OK: 一致する
	//h1 = human1(h3) // NG: 一致しない。（コンパイルエラー）
	
	// 無名関数の場合は直接代入できる
	var h4 struct {
		name string
		age int
	}
	h1 = h4
	
	fmt.Println(h1,h3)
}

