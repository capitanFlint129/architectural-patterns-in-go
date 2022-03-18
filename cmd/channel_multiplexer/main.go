package main

import (
	"context"
	"fmt"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/channel_multiplexer"
)

func main() {
	sig := func(data channel_multiplexer.ChannelDataStruct) <-chan channel_multiplexer.ChannelDataStruct {
		c := make(chan channel_multiplexer.ChannelDataStruct)
		go func() {
			defer close(c)
			c <- data
			time.Sleep(data.After)
		}()
		return c
	}

	start := time.Now()
	ctx := context.Background()
	multiplexedChannel := channel_multiplexer.OrRecursion(
		ctx,
		sig(channel_multiplexer.ChannelDataStruct{
			After: 1 * time.Second,
			Field: 100,
		}),
		sig(channel_multiplexer.ChannelDataStruct{
			After: 1 * time.Second,
			Field: 100,
		}),
	)

	for data := range multiplexedChannel {
		fmt.Println(data)
	}
	fmt.Printf("done after %v", time.Since(start))
}
