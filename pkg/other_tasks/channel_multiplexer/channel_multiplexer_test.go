package channel_multiplexer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type inputData struct {
	channelsData        []ChannelDataStruct
	funcForMultiplexing func(data ChannelDataStruct) <-chan ChannelDataStruct
}

type expectedResult struct {
	solution string
	error    error
	times    int
}

var (
	manyChannelsTestCaseName = "Many channels multiplexing"
	oneChannelTestCaseName   = "One channel"
)

func Test_OrRecursion(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: manyChannelsTestCaseName,
			inputData: inputData{
				channelsData: []ChannelDataStruct{
					{After: 1 * time.Nanosecond, Field: 100},
					{After: 2 * time.Nanosecond, Field: 200},
					{After: 3 * time.Nanosecond, Field: 300},
				},
				funcForMultiplexing: sendDataThenSleepAndClose,
			},
		},
		{
			testCaseName: oneChannelTestCaseName,
			inputData: inputData{
				channelsData: []ChannelDataStruct{
					{After: 1, Field: 100},
				},
				funcForMultiplexing: sendDataThenSleepAndClose,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			// Run funcForMultiplexing and create slice of channels that will be multiplexed
			channels := make([]<-chan ChannelDataStruct, 0)
			for _, data := range testData.inputData.channelsData {
				channels = append(channels, testData.inputData.funcForMultiplexing(data))
			}
			ctx := context.Background()
			// Multiplex channels
			multiplexedChannelCreator := NewMultiplexedChannelCreator(OrInner)
			ctxWithCancel, cancel := context.WithCancel(ctx)
			multiplexedChannel := multiplexedChannelCreator.GetMultiplexedChannel(
				&ContextWithCancel{
					Ctx:    ctxWithCancel,
					Cancel: cancel,
				},
				channels...,
			)
			// Close multiplexed channel after one of channels closed
			go func() {
				defer close(multiplexedChannel)
				<-ctxWithCancel.Done()
			}()
			// Collect data from channels through multiplexedChannel
			dataFromMultiplexedChannel := make([]ChannelDataStruct, 0)
			for data := range multiplexedChannel {
				dataFromMultiplexedChannel = append(dataFromMultiplexedChannel, data)
			}

			channelsIsClosedSlice := make([]bool, 0, len(channels))
			for _, channel := range channels {
				_, ok := <-channel
				channelsIsClosedSlice = append(channelsIsClosedSlice, !ok)
			}

			assert.EqualValues(t, getMapFromSlice(testData.inputData.channelsData), getMapFromSlice(dataFromMultiplexedChannel))
			assert.True(t, all(channelsIsClosedSlice))
		})
	}
}

func sendDataThenSleepAndClose(data ChannelDataStruct) <-chan ChannelDataStruct {
	c := make(chan ChannelDataStruct)
	go func() {
		defer close(c)
		c <- data
		time.Sleep(data.After)
	}()
	return c
}

func getMapFromSlice(slice []ChannelDataStruct) map[ChannelDataStruct]bool {
	mapFromSlice := make(map[ChannelDataStruct]bool)

	for _, data := range slice {
		mapFromSlice[data] = true
	}
	return mapFromSlice
}

func all(boolSlice []bool) bool {
	for _, value := range boolSlice {
		if !value {
			return false
		}
	}
	return true
}
