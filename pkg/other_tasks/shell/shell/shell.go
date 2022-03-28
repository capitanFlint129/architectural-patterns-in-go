package shell

import (
	"context"
	"sync"
)

type receiver interface {
	StartReceive(ctx context.Context, wg *sync.WaitGroup)
}

type processor interface {
	StartProcessing(ctx context.Context, wg *sync.WaitGroup)
}

type responder interface {
	StartRespond(ctx context.Context, wg *sync.WaitGroup)
}

type Shell interface {
	Run(ctx context.Context)
}

type shell struct {
	receiver  receiver
	processor processor
	responder responder
}

func (s *shell) Run(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	s.receiver.StartReceive(ctx, &wg)
	wg.Add(1)
	s.processor.StartProcessing(ctx, &wg)
	wg.Add(1)
	s.responder.StartRespond(ctx, &wg)

	wg.Wait()
}

func NewShell(receiver receiver, processor processor, responder responder) Shell {
	return &shell{
		receiver:  receiver,
		processor: processor,
		responder: responder,
	}
}
