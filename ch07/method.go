package main

import "fmt"

type Person struct {
	name string
	age int
}
// メソッド
func (p Person) toString() string {
	return fmt.Sprintf("%s: %d歳", p.name, p.age)
}

func main() {

	person := Person {
		name: "test",
		age: 30,
	}
	// メソッドの実行
	fmt.Println(person.toString())
}