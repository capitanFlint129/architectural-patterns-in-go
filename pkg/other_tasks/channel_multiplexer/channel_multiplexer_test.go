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
	noChannelsTestCaseName   = "No channels"
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
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			// Run funcForMultiplexing and create slice of channels that will be multiplexed
			channels := make([]<-chan ChannelDataStruct, 0)
			for _, data := range testData.inputData.channelsData {
				channels = append(channels, testData.inputData.funcForMultiplexing(data))
			}
			// Multiplex channels
			ctx := context.Background()
			multiplexedChannel := OrRecursion(ctx, channels...)
			// Collect data from channels through multiplexedChannel
			dataFromMultiplexedChannel := make([]ChannelDataStruct, 0)
			for data := range multiplexedChannel {
				dataFromMultiplexedChannel = append(dataFromMultiplexedChannel, data)
			}

			channelsStates := make([]bool, len(channels))
			for _, channel := range channels {
				_, ok := <-channel
				channelsStates = append(channelsStates, ok)
			}

			assert.EqualValues(t, getMapFromSlice(dataFromMultiplexedChannel), getMapFromSlice(testData.inputData.channelsData))
			assert.True(t, all(channelsStates))
		})
	}
}

func Test_OrRecursionNoChannels(t *testing.T) {
	for _, testData := range []struct {
		testCaseName string
	}{
		{
			testCaseName: noChannelsTestCaseName,
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			ctx := context.Background()
			multiplexedChannel := OrRecursion(ctx)
			assert.Nil(t, multiplexedChannel)
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
