package channel_multiplexer

import (
	"reflect"
	"time"
)

// ChannelDataStruct - data structure
type ChannelDataStruct struct {
	After time.Duration
	Field int
}

// OrRecursion accepts an arbitrary number of channels and returns
// one channel that returns the events any of the internal
func OrRecursion(channels ...<-chan ChannelDataStruct) <-chan ChannelDataStruct {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	multiplexedChannel := make(chan ChannelDataStruct)
	done := orInnerRecursion(multiplexedChannel, channels...)

	go func() {
		defer close(multiplexedChannel)
		<-done
	}()
	return multiplexedChannel
}

func orInnerRecursion(multiplexedChannel chan ChannelDataStruct, channels ...<-chan ChannelDataStruct) <-chan ChannelDataStruct {
	if len(channels) == 1 {
		return channels[0]
	}

	orDone := make(chan ChannelDataStruct)
	go func() {
		defer close(orDone)
	forLoop:
		for {
			var data ChannelDataStruct
			var ok bool
			select {
			case data, ok = <-channels[0]:
			case data, ok = <-channels[1]:
			case data, ok = <-orInnerRecursion(multiplexedChannel, append(channels[2:], orDone)...):
			}
			if !ok {
				break forLoop
			}
			multiplexedChannel <- data
		}
	}()
	return orDone
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
