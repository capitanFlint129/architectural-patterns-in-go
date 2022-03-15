package channel_multiplexer

import (
	"fmt"
	"reflect"
)

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

func OrRecursion(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	multiplexedChannel := make(chan interface{})
	done := OrInnerRecursion(multiplexedChannel, channels...)

	go func() {
	forLoop:
		for {
			select {
			case data := <-multiplexedChannel:
				fmt.Println(data)
			case <-done:
				close(multiplexedChannel)
				break forLoop
			}
		}
	}()
	return multiplexedChannel
}

func OrInnerRecursion(multiplexedChannel chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 1 {
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
	forLoop:
		for {
			var data interface{}
			var ok bool
			select {
			case data, ok = <-channels[0]:
			case data, ok = <-channels[1]:
			case data, ok = <-OrInnerRecursion(multiplexedChannel, append(channels[2:], orDone)...):
			}
			if !ok {
				break forLoop
			}
			multiplexedChannel <- data
		}
	}()
	return orDone
}
