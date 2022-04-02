package parser

import "strings"

type Parser interface {
	Parse(command string) []string
}

type parser struct {
	delimiter string
}

func (p *parser) Parse(command string) []string {
	if command == "" {
		return make([]string, 0)
	}
	commands := strings.Split(command, p.delimiter)
	for i := range commands {
		commands[i] = strings.TrimSpace(commands[i])
	}
	return commands
}

func NewParser(delimiter string) Parser {
	return &parser{delimiter: delimiter}
}
