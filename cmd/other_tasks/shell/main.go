package main

import (
	"bufio"
	"context"
	"os"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/command"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/parser"
	processor "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/processor"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/reciever"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/responder"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/shell"
)

const (
	pipeDelimiter    = "|"
	commandDelimiter = " "
)

var commandsCreatorsMap = map[string]func(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) command.Command{
	"cd":   command.NewCdCommand,
	"pwd":  command.NewPwdCommand,
	"echo": command.NewEchoCommand,
	"kill": command.NewKillCommand,
	"ps":   command.NewPsCommand,
	"fork": command.NewForkCommand,
	"exec": command.NewExecCommand,
}

func main() {
	errorChannel := make(chan error)
	receiverOutputChannel := make(chan string)
	responderInputChannel := make(chan string)

	scanner := bufio.NewScanner(os.Stdin)
	newReceiver := reciever.NewReceiver(scanner, receiverOutputChannel, errorChannel)
	newResponder := responder.NewResponder(os.Stdin, os.Stderr, responderInputChannel, errorChannel)

	newPipeParser := parser.NewParser(pipeDelimiter)
	newCommandParser := parser.NewParser(commandDelimiter)
	newProcessor := processor.NewProcessor(newPipeParser, newCommandParser, receiverOutputChannel, responderInputChannel, errorChannel, commandsCreatorsMap)
	newShell := shell.NewShell(newReceiver, newProcessor, newResponder)

	mainCtx := context.Background()
	ctx, _ := context.WithTimeout(mainCtx, 10*time.Minute)
	newShell.Run(ctx)

}
