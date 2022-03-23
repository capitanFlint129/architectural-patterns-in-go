package channel_multiplexer

import (
	"context"
	"reflect"
	"time"
)

// ChannelDataStruct - data structure
type ChannelDataStruct struct {
	After time.Duration
	Field int
}

type ContextWithCancel struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

// Or accepts an arbitrary number of channels and returns
// one channel that returns the events any of the internal
func Or(parent context.Context, multiplexedChannelCreator MultiplexedChannelCreator, channels ...<-chan ChannelDataStruct) <-chan ChannelDataStruct {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	ctx, cancel := context.WithCancel(parent)
	multiplexedChannel := multiplexedChannelCreator.GetMultiplexedChannel(&ContextWithCancel{
		Ctx:    ctx,
		Cancel: cancel,
	}, channels...)

	go func() {
		defer close(multiplexedChannel)
		<-ctx.Done()
	}()
	return multiplexedChannel
}

func OrInner(
	contextStruct *ContextWithCancel,
	multiplexedChannel chan ChannelDataStruct,
	channels ...<-chan ChannelDataStruct,
) {
	go func(contextStruct *ContextWithCancel) {
		defer contextStruct.Cancel()
		for {
			select {
			case data, ok := <-channels[0]:
				if !ok {
					return
				}
				multiplexedChannel <- data
			case _, ok := <-contextStruct.Ctx.Done():
				if !ok {
					return
				}
			}
		}
	}(contextStruct)
}

func orWithReflectSelect(channels ...<-chan interface{}) <-chan interface{} {
	cases := make([]reflect.SelectCase, len(channels))
	for i, channel := range channels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(channel)}
	}

	multiplexedChannel := make(chan interface{})
	go func() {
		for {
			_, _, isOpen := reflect.Select(cases)
			if isOpen == false {
				close(multiplexedChannel)
				return
			}
		}
	}()
	return multiplexedChannel
}
