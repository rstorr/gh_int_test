package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func test4() {
	f, err := os.Open("/tmp/test.txt") //ISSUE
	check(err)
	//defer f.Close()
	b := make([]byte, 5)
	n, err := f.Read(b)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(b))
}

func main() {

	f1, err := os.Open("/tmp/test.txt") //ISSUE
	check(err)
	//defer f1.Close()
	b1 := make([]byte, 5)
	n1, err := f1.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	f2, err := os.Open("/tmp/test.txt")
	check(err)
	//defer f2.Close()
	b2 := make([]byte, 5)
	n2, err := f2.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %s\n", n2, string(b2))
}
