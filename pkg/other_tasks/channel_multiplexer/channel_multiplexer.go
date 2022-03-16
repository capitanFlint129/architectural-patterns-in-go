package channel_multiplexer

import (
	"reflect"
	"time"
)

type ChannelDataStruct struct {
	After time.Duration
	Field int
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

func OrRecursion(channels ...<-chan ChannelDataStruct) <-chan ChannelDataStruct {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	multiplexedChannel := make(chan ChannelDataStruct)
	done := OrInnerRecursion(multiplexedChannel, channels...)

	go func() {
		defer close(multiplexedChannel)
		<-done
	}()
	return multiplexedChannel
}

func OrInnerRecursion(multiplexedChannel chan ChannelDataStruct, channels ...<-chan ChannelDataStruct) <-chan ChannelDataStruct {
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

//func isClosed(ch <-chan interface{}) bool {
//	//select {
//	//case <-ch:
//	//	return true
//	//default:
//	//}
//	//return false
//}
//
//func or(channels ...<-chan interface{}) <-chan interface{} {
//	multiplexedChannel := make(chan interface{})
//	done := make(chan interface{})
//	for _, channel := range channels {
//		go func(channel <-chan interface{}, done chan interface{}) {
//		forLoop:
//			for {
//				select {
//				case data, isOpen := <-channel:
//					if isOpen && !isClosed(multiplexedChannel) {
//						multiplexedChannel <- data
//					} else {
//						if !isClosed(done) {
//							close(done)
//						}
//						if !isClosed(multiplexedChannel) {
//							close(multiplexedChannel)
//						}
//					}
//				case _, isDone := <-done:
//					if isDone {
//						fmt.Println("Closed")
//						break forLoop
//					}
//				}
//			}
//		}(channel, done)
//	}
//	return multiplexedChannel
//}
