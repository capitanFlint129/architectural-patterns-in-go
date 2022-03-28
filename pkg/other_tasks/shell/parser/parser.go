package parser

import "strings"

type Parser interface {
	Parse(pipeCommand string) []string
}

type parser struct {
	delimiter string
}

func (p *parser) Parse(pipeCommand string) []string {
	pipeCommand = strings.TrimSuffix(pipeCommand, "\n")
	if pipeCommand == "" {
		return make([]string, 0)
	}
	commands := strings.Split(pipeCommand, p.delimiter)
	return commands
}

func NewParser(delimiter string) Parser {
	return &parser{delimiter: delimiter}
}
