package channel_multiplexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type inputData struct {
	channelsData []ChannelDataStruct
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
			sig := func(data ChannelDataStruct) <-chan ChannelDataStruct {
				c := make(chan ChannelDataStruct)
				go func() {
					defer close(c)
					c <- data
					time.Sleep(data.After)
				}()
				return c
			}

			channels := make([]<-chan ChannelDataStruct, 0)
			for _, data := range testData.inputData.channelsData {
				channels = append(channels, sig(data))
			}
			multiplexedChannel := OrRecursion(channels...)

			dataFromMultiplexedChannel := make([]ChannelDataStruct, 0)
			for i := 0; i < len(channels); i++ {
				data, ok := <-multiplexedChannel
				assert.True(t, ok)
				dataFromMultiplexedChannel = append(dataFromMultiplexedChannel, data)
			}

			assert.EqualValues(t, getMapFromSlice(dataFromMultiplexedChannel), getMapFromSlice(testData.inputData.channelsData))
			_, isOpen := <-multiplexedChannel
			assert.True(t, !isOpen)
			for _, channel := range channels {
				_, isOpen = <-channel
				assert.True(t, !isOpen)
			}
		})
	}
}

func Test_OrRecursionNoChannles(t *testing.T) {
	for _, testData := range []struct {
		testCaseName string
	}{
		{
			testCaseName: noChannelsTestCaseName,
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			multiplexedChannel := OrRecursion()

			assert.Nil(t, multiplexedChannel)
		})
	}
}

func getMapFromSlice(slice []ChannelDataStruct) map[ChannelDataStruct]bool {
	mapFromSlice := make(map[ChannelDataStruct]bool)

	for _, data := range slice {
		mapFromSlice[data] = true
	}
	return mapFromSlice
}
