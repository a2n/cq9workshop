package main

import "log"

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	f4()
}
