package main

import (
	"errors"
	"fmt"
)

func main() {
	e0 := errors.New("e0")
	e1 := fmt.Errorf("e1: %w", e0)
	fmt.Printf("e0 '%s', e1 '%s'\n", e0.Error(), e1.Error())
	fmt.Printf("e0 is e1 %t\n", errors.Is(e0, e1))
	fmt.Printf("e0 is e1 %t\n", errors.Is(e0, errors.Unwrap(e1)))
}
