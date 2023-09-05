package main

import (
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) {
	if i % 2 != 0 {
		// 文字列を返す場合は、errors.NewでOK
		// return 0, errors.New("偶数ではありません")
		// fmt.Errorfでフォーマットも使用できる
		return 0, fmt.Errorf("%dは偶数ではありません", i)
	}
	return i * 2, nil
}

func main() {
	i := 17
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%dの2倍は %d\n", i, double)
}