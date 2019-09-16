package main

import (
	"fmt"
)

func main() {
	n := 1
	e := isEven(n)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("%d is a even number\n", n)
	}

	n = 2
	e = isEven(n)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("%d is a even number\n", n)
	}
}

func isEven(n int) error {
	if n%2 == 1 {
		return fmt.Errorf("%d is a odd number", n)
	}
	return nil
}
