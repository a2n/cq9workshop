package main

import (
	"crypto/rand"
	"errors"
	"log"
	"math"
	"math/big"
)

func main() {
	log.SetFlags(log.Lshortfile)
	handler()
}

func handler() {
	pkg()
}

func pkg() {
	pkgA()
	pkgB()
}

func pkgA() {
	e := db()
	if e != nil {
		log.Printf("pkgA, %s", e.Error())
		return
	}
}

func pkgB() {
	e := db()
	if e != nil {
		log.Printf("pkgB, %s", e.Error())
		return
	}
}

func db() error {
	if randEven() {
		return errors.New("something wrong")
	}
	return nil
}

func randEven() bool {
	max := big.NewInt(math.MaxInt64)
	r, e := rand.Int(rand.Reader, max)
	if e != nil {
		return false
	}
	if r.Int64()%2 == 0 {
		return true
	}
	return false
}
