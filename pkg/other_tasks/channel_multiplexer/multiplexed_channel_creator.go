package channel_multiplexer

type channelCreatorFunc func(
	contextStruct *ContextWithCancel,
	multiplexedChannel chan ChannelDataStruct,
	channels ...<-chan ChannelDataStruct,
)

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
		multiplexedChannel = m.GetMultiplexedChannel(contextStruct, channels[1:]...)
		m.creator(contextStruct, multiplexedChannel, channels...)
	}
	return multiplexedChannel
}

func NewMultiplexedChannelCreator(creator channelCreatorFunc) MultiplexedChannelCreator {
	return &multiplexedChannelCreator{
		creator: creator,
	}
}
