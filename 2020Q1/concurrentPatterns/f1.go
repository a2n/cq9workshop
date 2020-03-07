package main

import (
	"log"
	"time"
)

func f1() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			log.Print("tick")
		case <-timeout:
			log.Print("time is up, quit")
			return
		}
	}
}
