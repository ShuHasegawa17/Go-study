package main

import "fmt"

// 具象型
type LogicProvider struct{}

// 具象型のメソッド
func(lp LogicProvider) Process(data string) string {
	fmt.Println("具象型のメソッド:", data)
	return data
}

// インターフェース
type Logic interface {
	// インターフェースのメソッド定義
	Process(data string) string
	execute()int
}

// 利用する型
type Client struct{
	L Logic //インターフェースをメンバで定義
}

// 利用する型のメソッド
func (c Client) Program() {
	data := "use-client"
	// インターフェースのメソッド呼び出し
	c.L.Process(data)
}

func main() {
	// 実行時に、利用する側が、具象型を設定する
	c := Client {
		//L: LogicProvider{},
	}
	c.Program()
}