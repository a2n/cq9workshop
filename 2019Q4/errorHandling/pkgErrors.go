package main

import (
	"crypto/rand"
	goErrors "errors"
	"log"
	"math"
	"math/big"

	"github.com/pkg/errors"
)

func main() {
	log.SetFlags(log.Lshortfile)
	e := handler()
	if e != nil {
		log.Printf("%+v", e)
		return
	}
}

func handler() error {
	e := pkg()
	if e != nil {
		return e
	}
	return nil
}

func pkg() error {
	e := pkgA()
	if e != nil {
		return e
	}

	e = pkgB()
	if e != nil {
		return e
	}

	return nil
}

func pkgA() error {
	e := db()
	if e != nil {
		return errors.Wrap(e, "pkgA")
	}
	return nil
}

func pkgB() error {
	e := db()
	if e != nil {
		return errors.WithStack(e)
	}
	return nil
}

func db() error {
	if randEven() {
		return goErrors.New("something wrong")
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
