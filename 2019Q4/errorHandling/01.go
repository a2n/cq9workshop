package main

import (
	"crypto/rand"
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
	db()
}

func pkgB() {
	db()
}

func db() {
	if randEven() {
		log.Print("something wrong")
		return
	}
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
