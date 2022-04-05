package responder

import (
	"context"
	"fmt"
	"io"
	"sync"
)

type Responder interface {
	StartRespond(ctx context.Context, wg *sync.WaitGroup)
}

type responder struct {
	outputWriter io.Writer
	errorWriter  io.Writer
	inputChannel <-chan string
	errorChannel <-chan error
}

func (r *responder) StartRespond(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			select {
			case output := <-r.inputChannel:
				if _, err := fmt.Fprintln(r.outputWriter, output); err != nil {
					fmt.Fprintln(r.errorWriter, err)
				}
			case err := <-r.errorChannel:
				fmt.Fprintln(r.errorWriter, err)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func NewResponder(outputWriter io.Writer, errorWriter io.Writer, inputChannel <-chan string, errorChannel <-chan error) Responder {
	return &responder{
		outputWriter: outputWriter,
		errorWriter:  errorWriter,
		inputChannel: inputChannel,
		errorChannel: errorChannel,
	}
}
