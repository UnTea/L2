package main

import (
	"testing"
	"time"
)

const realOut = 1

func TestOrChannel(t *testing.T) {
	signal := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()

	<-Or(
		signal(2*time.Hour),
		signal(5*time.Minute),
		signal(1*time.Second),
		signal(1*time.Hour),
		signal(2*time.Second),
	)

	timeAfter := int(time.Since(start).Seconds())

	if timeAfter != realOut {
		t.Errorf("%v!=%v\nShould: %d\nGot: %d\n", realOut, timeAfter, realOut, timeAfter)
	}
}
