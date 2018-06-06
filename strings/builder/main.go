// Package main provides ...
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	fmt.Printf("初始状态 = %+v\n", b)
	for i := 0; i < 3; i++ {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
	fmt.Printf("b.Len() = %+v\n", b.Len())
	fmt.Printf("b.Len() == len(b.String()) = %+v\n", b.Len() == len(b.String()))
	// Grow
	// b.Grow(30)
	fmt.Printf("添加元素后 = %+v\n", b)
	b.Reset()
	fmt.Printf("Rest后 = %+v\n", b)
}
