package responder

import (
	"context"
	"fmt"
	"os"
	"sync"
)

type Responder interface {
	StartRespond(ctx context.Context, wg *sync.WaitGroup)
}

type responder struct {
	file         *os.File
	errFile      *os.File
	inputChannel <-chan string
	errorChannel <-chan error
}

func (r *responder) StartRespond(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			select {
			case output := <-r.inputChannel:
				_, err := fmt.Fprintln(r.file, output)
				if err != nil {
					fmt.Fprintln(r.errFile, err)
				}
			case err := <-r.errorChannel:
				fmt.Fprintln(r.errFile, err)
			case <-ctx.Done():
				break
			}
		}
	}()
}

func NewResponder(file *os.File, errFile *os.File, inputChannel <-chan string, errorChannel <-chan error) Responder {
	return &responder{
		file:         file,
		errFile:      errFile,
		inputChannel: inputChannel,
		errorChannel: errorChannel,
	}
}
