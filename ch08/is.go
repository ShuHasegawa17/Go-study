package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker (name string) error{
	f,err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChekcer: %w", err); // %wのラップ
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("notfound.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("not exisit")
		}
	}
}