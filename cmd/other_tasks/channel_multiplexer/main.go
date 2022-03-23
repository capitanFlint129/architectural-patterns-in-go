package main

import (
	"context"
	"fmt"
	channel_multiplexer "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/channel_multiplexer"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	multiplexedChannelCreator := channel_multiplexer.NewMultiplexedChannelCreator(channel_multiplexer.OrInner)

	ctxWithCancel, cancel := context.WithCancel(ctx)

	multiplexedChannel := multiplexedChannelCreator.GetMultiplexedChannel(
		&channel_multiplexer.ContextWithCancel{
			Ctx:    ctxWithCancel,
			Cancel: cancel,
		},
		sig(channel_multiplexer.ChannelDataStruct{
			After: 1 * time.Millisecond,
			Field: 100,
		}),
		sig(channel_multiplexer.ChannelDataStruct{
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

func sig(data channel_multiplexer.ChannelDataStruct) <-chan channel_multiplexer.ChannelDataStruct {
	c := make(chan channel_multiplexer.ChannelDataStruct)
	go func() {
		defer close(c)
		c <- data
		time.Sleep(data.After)
	}()
	return c
}
