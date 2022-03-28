package reciever

import (
	"bufio"
	"context"
	"sync"
)

type Receiver interface {
	StartReceive(ctx context.Context, wg *sync.WaitGroup)
}

type receiver struct {
	scanner       *bufio.Scanner
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (r *receiver) StartReceive(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for r.scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
			}
			r.outputChannel <- r.scanner.Text()
		}
		if r.scanner.Err() != nil {
			r.errorChannel <- r.scanner.Err()
		}
	}()
}

func NewReceiver(scanner *bufio.Scanner, outputChannel chan<- string, errorChannel chan<- error) Receiver {
	return &receiver{
		scanner:       scanner,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
