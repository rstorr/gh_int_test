package main

import (
	"fmt"
	"os"
)

func test2() {
	f, err := os.Open("/tmp/test.txt") //ISSUE
	check(err)
	//defer f.Close()
	b := make([]byte, 5)
	n, err := f.Read(b)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(b))
}

func test3() {
	f, err := os.Open("/tmp/test.txt") //ISSUE
	check(err)
	//defer f.Close()
	b := make([]byte, 5)
	n, err := f.Read(b)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(b))
}
