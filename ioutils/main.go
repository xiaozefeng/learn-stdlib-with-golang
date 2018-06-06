// Package main provides ...
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// r := strings.NewReader("Hello World, 你好世界")
	f, err := os.Open("../README.md")

	// r只要实现了io.Reader接口就可以读取
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b) = %+v\n", string(b))

	// 读取文件夹
	files, err := ioutil.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range files {
		fmt.Printf("v.Name() = %+v\n", v.Name())
	}

	// 直接读取文件
	content, err := ioutil.ReadFile("../README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(content) = %+v\n", string(content))

}
