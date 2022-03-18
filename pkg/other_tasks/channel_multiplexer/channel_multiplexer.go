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

type contextWithCancel struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// OrRecursion accepts an arbitrary number of channels and returns
// one channel that returns the events any of the internal
func OrRecursion(parent context.Context, channels ...<-chan ChannelDataStruct) <-chan ChannelDataStruct {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	ctx, cancel := context.WithCancel(parent)
	multiplexedChannel := orInnerRecursion(&contextWithCancel{
		ctx:    ctx,
		cancel: cancel,
	}, channels...)

	go func() {
		defer close(multiplexedChannel)
		<-ctx.Done()
	}()
	return multiplexedChannel
}

func orInnerRecursion(
	contextStruct *contextWithCancel,
	channels ...<-chan ChannelDataStruct,
) chan ChannelDataStruct {
	var multiplexedChannel chan ChannelDataStruct

	// TODO вопрос: получается слишком нагроможденно, отдельная функция будет вызывать orInnerRecursion (рекурсия станет неявной)
	// TODO вынести if в пакет
	if len(channels) == 1 {
		multiplexedChannel = make(chan ChannelDataStruct)
	} else {
		multiplexedChannel = orInnerRecursion(contextStruct, channels[1:]...)
	}

	go func(contextStruct *contextWithCancel) {
		defer contextStruct.cancel()
		for {
			select {
			case data, ok := <-channels[0]:
				if !ok {
					return
				}
				multiplexedChannel <- data
			case _, ok := <-contextStruct.ctx.Done():
				if !ok {
					return
				}
			}
		}
	}(contextStruct)
	return multiplexedChannel
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
