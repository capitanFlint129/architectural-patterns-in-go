package command

import "sync"

type Command = interface {
	Execute(wg *sync.WaitGroup)
}
