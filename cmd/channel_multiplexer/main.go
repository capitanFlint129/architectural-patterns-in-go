package main

import (
	"context"
	"fmt"
	channel_multiplexer2 "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/channel_multiplexer"
	"time"
)

func main() {
	sig := func(data channel_multiplexer2.ChannelDataStruct) <-chan channel_multiplexer2.ChannelDataStruct {
		c := make(chan channel_multiplexer2.ChannelDataStruct)
		go func() {
			defer close(c)
			c <- data
			time.Sleep(data.After)
		}()
		return c
	}

	start := time.Now()
	ctx := context.Background()
	multiplexedChannelCreator := channel_multiplexer2.NewMultiplexedChannelCreator(channel_multiplexer2.OrInner)

	ctxWithCancel, cancel := context.WithCancel(ctx)

	multiplexedChannel := multiplexedChannelCreator.GetMultiplexedChannel(
		&channel_multiplexer2.ContextWithCancel{
			Ctx:    ctxWithCancel,
			Cancel: cancel,
		},
		sig(channel_multiplexer2.ChannelDataStruct{
			After: 1 * time.Millisecond,
			Field: 100,
		}),
		sig(channel_multiplexer2.ChannelDataStruct{
			After: 1 * time.Second,
			Field: 200,
		}),
	)

	go func() {
		defer close(multiplexedChannel)
		<-ctxWithCancel.Done()
	}()
	for data := range multiplexedChannel {
		fmt.Println(data)
	}

	fmt.Printf("done after %v", time.Since(start))
}
