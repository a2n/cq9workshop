package main

import (
	"context"
	"log"
	"time"
)

func f4() {
	f4Values()
}

func f4Values() {
	type valueKey string
	var k valueKey = "key"
	f := func(ctx context.Context) {
		v := ctx.Value(k)
		if v == nil {
			log.Print("nil value")
			return
		}
		val, ok := v.(string)
		if ok == false {
			log.Print("not string type")
			return
		}
		log.Printf("value: '%s'", val)
	}

	ctx := context.WithValue(context.Background(), k, "it's me")
	f(ctx)
}

func f4Timeout() {
	f := func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				log.Print("tick")
			case <-ctx.Done():
				log.Print("time is up, quit")
				return
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	f(ctx)
}
