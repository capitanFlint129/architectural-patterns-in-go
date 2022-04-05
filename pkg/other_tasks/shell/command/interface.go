package command

import (
	"context"
	"sync"
)

// Command for shell
type Command = interface {
	// Execute - executes command
	Execute(ctx context.Context, wg *sync.WaitGroup)
	SetArgs(args []string) error
}
