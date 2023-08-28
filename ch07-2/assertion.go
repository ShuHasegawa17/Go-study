package main

import (
	"fmt"
	"os"
)

type MyInt int
func main() {
	var i any
	var mine MyInt = 20
	i = mine;
	i2 := i.(MyInt) // 型アサーション
	i3 := i.(string) // 型アサーションが違うのでパニック
	i4 := i.(int) // 基底型は同じだが、実際の型はMyIntなのでパニック

	i5,ok := i.(int) // カンマokイディオムで確認する
	if !ok {
		err := fmt.Errorf("iの型（値:%v）が想定外",i)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(i2,i3,i4,i5);
}

func doTypeSwitch(i any) any  {
	switch i := i.(type) {
	case nil:
		// ...
	case int:
		// ...
	case string:
		fmt.Println(i)
	}
	return i
}