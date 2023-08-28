package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total int
	updated time.Time
}

// ポインタ型レシーバ
func (c *Counter) Incremament() {
	c.total++
	c.updated = time.Now()
}

func (c Counter) toString() {
	fmt.Print("total:", c.total, ",time:", c.updated)
}

// OK
func doUpdate(c *Counter) {
	c.Incremament()
	fmt.Println("OK:", c.total ,c.updated)
}

// NG:値レシーバ
func doUpdate2(c Counter) {
	c.Incremament()
	fmt.Println("NG:", c.total ,c.updated)
}

func main() {
	var c Counter
	c.Incremament()
	doUpdate(&c)
	c.toString()
	doUpdate2(c)
	c.toString()
}