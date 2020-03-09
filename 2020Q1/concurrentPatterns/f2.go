package main

import (
	"log"
	"time"
)

func f2() {
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("blog.google")),
		Subscribe(Fetch("developers.googleblog.com")),
	)

	time.AfterFunc(3*time.Second, func() {
		log.Print("Closed: ", merged.Close())
	})

	for i := range merged.Updates() {
		log.Print(i.Title, i.Channel)
	}
}

type Item struct {
	Title, Channel, GUID string
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

func Fetch(domain string) Fetcher {
	return nil
}

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	go s.loop()
	return s
}

func Merge(subs ...Subscription) Subscription {
	return nil
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	closed  bool
	err     error

	closing chan chan error
}

func (s *sub) loop() {
	var err error
	for {
		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		}

		if s.closed {
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err
			time.Sleep(10 * time.Second)
			continue
		}
		for _, item := range items {
			s.updates <- item
		}
		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now))
		}
	}
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}
