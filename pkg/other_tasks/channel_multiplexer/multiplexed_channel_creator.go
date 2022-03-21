package channel_multiplexer

type channelCreatorFunc func(
	contextStruct *ContextWithCancel,
	multiplexedChannelCreator MultiplexedChannelCreator,
	channels ...<-chan ChannelDataStruct,
) chan ChannelDataStruct

type MultiplexedChannelCreator interface {
	GetMultiplexedChannel(
		contextStruct *ContextWithCancel,
		channels ...<-chan ChannelDataStruct,
	) chan ChannelDataStruct
}

type multiplexedChannelCreator struct {
	creator channelCreatorFunc
}

func (m *multiplexedChannelCreator) GetMultiplexedChannel(
	contextStruct *ContextWithCancel,
	channels ...<-chan ChannelDataStruct,
) chan ChannelDataStruct {
	var multiplexedChannel chan ChannelDataStruct
	if len(channels) == 0 {
		multiplexedChannel = make(chan ChannelDataStruct)
	} else {
		multiplexedChannel = m.creator(contextStruct, m, channels...)
	}
	return multiplexedChannel
}

func NewMultiplexedChannelCreator(creator channelCreatorFunc) MultiplexedChannelCreator {
	return &multiplexedChannelCreator{
		creator: creator,
	}
}
