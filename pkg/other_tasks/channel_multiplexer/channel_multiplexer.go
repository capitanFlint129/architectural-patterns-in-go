package main

import (
	"fmt"
	"reflect"
	"time"
)

func orWithReflectSelect(channels ...<-chan interface{}) <-chan interface{} {
	cases := make([]reflect.SelectCase, len(channels))
	for i, channel := range channels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(channel)}
	}

	multiplexedChannel := make(chan interface{})
	go func() {
		for {
			_, _, is_open := reflect.Select(cases)
			if is_open == false {
				close(multiplexedChannel)
				return
			}
		}
	}()
	return multiplexedChannel
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	multiplexedChannel := make(chan interface{})
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			_, is_open := <-channel
			if is_open == false {
				close(multiplexedChannel)
				return
			}
		}(channel)
	}
	return multiplexedChannel
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("after %v", time.Since(start))
}
