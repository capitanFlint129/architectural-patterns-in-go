package main

import (
	"fmt"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/channel_multiplexer"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			c <- struct{ field int }{100}
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	multiplexedChannel := channel_multiplexer.OrRecursion(
		sig(1 * time.Second),
	)

	for data := range multiplexedChannel {
		fmt.Println(data)
	}
	fmt.Printf("done after %v", time.Since(start))
}
