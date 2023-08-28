package main

import "fmt"


type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Closer() error
}
type ReadCloser interface {
	Reader
	Closer
}

func main() {

	var s *string
	fmt.Println(s == nil) // true
	var i any // interfase{}でも同じ
	fmt.Println(i == nil) // true
	i = s // 型のポインタを渡している
	fmt.Println(i == nil) // false: 値のポインタはnil、型のポインタは非nil

}
